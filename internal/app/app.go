package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/shamank/ai-marketplace-stats-service/internal/config"
	"github.com/shamank/ai-marketplace-stats-service/internal/logger"
	"github.com/shamank/ai-marketplace-stats-service/internal/repository"
	"github.com/shamank/ai-marketplace-stats-service/internal/service"
	"log/slog"
)

type App struct {
	cfg        *config.Config
	log        *slog.Logger
	gRPCServer *GRPCServer
}

func NewApp(cfg *config.Config) *App {
	log := logger.InitLogger()

	fmt.Printf("config: %+v\n", cfg)

	dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {

		log.Error("failed to connect to database", err)
		return nil
	}

	repos := repository.NewRepositories(conn)

	services := service.NewService(log, repos)

	grpcServer := NewGRPCServer(log, services, cfg.GRPC.Port)

	return &App{
		cfg:        cfg,
		log:        log,
		gRPCServer: grpcServer,
	}
}

func (a *App) Run() error {

	a.log.Info("Starting application")

	if err := a.gRPCServer.Run(); err != nil {
		return err
	}

	return nil

}

func (a *App) Stop() {

	a.log.Info("Stopping application")
	a.gRPCServer.Stop()
}
