package main

import (
	"context"
	"flag"
	"fmt"
	grpRepo "github.com/isd-sgcu/rnkm65-backend/src/app/repository/group"
	ur "github.com/isd-sgcu/rnkm65-backend/src/app/repository/user"
	fSrv "github.com/isd-sgcu/rnkm65-backend/src/app/service/file"
	grpService "github.com/isd-sgcu/rnkm65-backend/src/app/service/group"
	us "github.com/isd-sgcu/rnkm65-backend/src/app/service/user"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	"github.com/isd-sgcu/rnkm65-backend/src/database"
	seed "github.com/isd-sgcu/rnkm65-backend/src/database/seeds"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
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
				log.Fatal().
					Str("service", "seeder").
					Msg("Not found seed")
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

		log.Info().
			Str("service", "graceful shutdown").
			Msgf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Error().
				Str("service", "graceful shutdown").
				Msgf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
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

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Error().
						Str("service", "graceful shutdown").
						Err(err).
						Msgf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("%v was shutdown gracefully", innerKey)
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
		log.Fatal().
			Err(err).
			Str("service", "auth").
			Msg("Failed to start service")
	}

	db, err := database.InitDatabase(&conf.Database)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "auth").
			Msg("Failed to start service")
	}

	cache, err := database.InitRedisConnect(&conf.Redis)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "auth").
			Msg("Failed to start service")
	}

	handleArgs(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "auth").
			Msg("Failed to start service")
	}

	fileConn, err := grpc.Dial(conf.Service.File, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "rnkm-file").
			Msg("Cannot connect to service")
	}

	grpcServer := grpc.NewServer()

	fileClient := proto.NewFileServiceClient(fileConn)
	fileSrv := fSrv.NewService(fileClient)

	usrRepo := ur.NewRepository(db)
	usrSvc := us.NewService(usrRepo, fileSrv)

	groupRepo := grpRepo.NewRepository(db)
	grpSvc := grpService.NewService(groupRepo, usrRepo)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	proto.RegisterUserServiceServer(grpcServer, usrSvc)
	proto.RegisterGroupServiceServer(grpcServer, grpSvc)
	reflection.Register(grpcServer)
	go func() {
		log.Info().
			Str("service", "auth").
			Msgf("RNKM65 backend starting at port %v", conf.App.Port)

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal().
				Err(err).
				Str("service", "auth").
				Msg("Failed to start service")
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
	log.Info().
		Str("service", "auth").
		Msg("Closing the listener")
	lis.Close()
	log.Info().
		Str("service", "auth").
		Msg("End of Program")
}
