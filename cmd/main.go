package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
	"github.com/isd-sgcu/rpkm66-backend/database"
	seed "github.com/isd-sgcu/rpkm66-backend/database/seeds"
	baan_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/baan/v1"
	checkin_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/checkin/v1"
	event_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/event/v1"
	group_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/group/v1"
	user_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/user/v1"
	bRepo "github.com/isd-sgcu/rpkm66-backend/internal/repository/baan"
	bgsRepo "github.com/isd-sgcu/rpkm66-backend/internal/repository/baan-group-selection"
	"github.com/isd-sgcu/rpkm66-backend/internal/repository/cache"
	cir "github.com/isd-sgcu/rpkm66-backend/internal/repository/checkin"
	evtRepo "github.com/isd-sgcu/rpkm66-backend/internal/repository/event"
	grpRepo "github.com/isd-sgcu/rpkm66-backend/internal/repository/group"
	ur "github.com/isd-sgcu/rpkm66-backend/internal/repository/user"
	bSrv "github.com/isd-sgcu/rpkm66-backend/internal/service/baan"
	csr "github.com/isd-sgcu/rpkm66-backend/internal/service/checkin"
	evtService "github.com/isd-sgcu/rpkm66-backend/internal/service/event"
	fSrv "github.com/isd-sgcu/rpkm66-backend/internal/service/file"
	grpService "github.com/isd-sgcu/rpkm66-backend/internal/service/group"
	us "github.com/isd-sgcu/rpkm66-backend/internal/service/user"
	file_proto "github.com/isd-sgcu/rpkm66-go-proto/rpkm66/file/file/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
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
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	db, err := database.InitDatabase(&conf.Database)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	cacheDB, err := database.InitRedisConnect(&conf.Redis)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	handleArgs(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
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

	cacheRepo := cache.NewRepository(cacheDB)

	fileClient := file_proto.NewFileServiceClient(fileConn)
	fileSrv := fSrv.NewService(fileClient)

	eventRepo := evtRepo.NewRepository(db)
	evtSvc := evtService.NewService(eventRepo)

	usrRepo := ur.NewRepository(db)
	usrSvc := us.NewService(usrRepo, fileSrv, eventRepo)

	ciRepo := cir.NewRepository(db)
	ciSvc := csr.NewService(ciRepo, cacheRepo, conf.App)

	baRepo := bRepo.NewRepository(db)
	baSrv := bSrv.NewService(baRepo, cacheRepo, conf.App)

	baGrpSetRepo := bgsRepo.NewRepository(db)

	groupRepo := grpRepo.NewRepository(db)
	grpSvc := grpService.NewService(groupRepo, usrRepo, baGrpSetRepo, fileSrv, cacheRepo, baRepo, conf.App)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	user_proto.RegisterUserServiceServer(grpcServer, usrSvc)
	checkin_proto.RegisterCheckinServiceServer(grpcServer, ciSvc)
	baan_proto.RegisterBaanServiceServer(grpcServer, baSrv)
	group_proto.RegisterGroupServiceServer(grpcServer, grpSvc)
	event_proto.RegisterEventServiceServer(grpcServer, evtSvc)

	reflection.Register(grpcServer)
	go func() {
		log.Info().
			Str("service", "backend").
			Msgf("rpkm66 backend starting at port %v", conf.App.Port)

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal().
				Err(err).
				Str("service", "backend").
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
			return cacheDB.Close()
		},
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	<-wait

	grpcServer.GracefulStop()
	log.Info().
		Str("service", "backend").
		Msg("Closing the listener")
	lis.Close()
	log.Info().
		Str("service", "backend").
		Msg("End of Program")
}
