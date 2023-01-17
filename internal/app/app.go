package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitlab.mapcard.pro/external-map-team/shop-service/config"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/controller/grpc"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/usecase"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/usecase/repo"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LogLevel)

	pg, err := postgres.New(*cfg)

	if err != nil {
		l.Fatalf(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	shopUseCase := usecase.New(
		l,
		repo.New(pg, l),
	)

	grpcServer := grpc.NewServer(l, shopUseCase)

	err = grpcServer.Start(cfg.GrpcPort)

	if err != nil {
		l.Fatalf(fmt.Errorf("app - NewApp - grpcServer start error: %w", err))
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Infof("app - Run - signal: " + s.String())
	}

}
