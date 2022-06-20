package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	"github.com/isd-sgcu/rnkm65-backend/src/database"
	seed "github.com/isd-sgcu/rnkm65-backend/src/database/seeds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func handleArgs(db *gorm.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			err := seed.Execute(db, args[1:]...)
			if err != nil {
				log.Fatalln("Not found seed")
			}
			os.Exit(0)
		}
	}
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Printf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Printf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}

	db, err := database.InitDatabase(&conf.Database)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
	}

	cache, err := database.InitRedisConnect(&conf.Redis)
	if err != nil {
		log.Fatal("Cannot connect to redis: ", err.Error())
	}

	handleArgs(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	go func() {
		log.Println(fmt.Sprintf("RNKM65 backend starting at port %v", conf.App.Port))

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalln("Failed to serve:", err)
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			sqlDb, err := db.DB()
			if err != nil {
				return err
			}
			return sqlDb.Close()
		},
		"cache": func(ctx context.Context) error {
			return cache.Close()
		},
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	<-wait

	grpcServer.GracefulStop()
	log.Println("Closing the listener")
	lis.Close()
	log.Println("End of Program")
}
