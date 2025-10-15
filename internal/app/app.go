package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/services/auth"
	"time"
)

type (
	App struct {
		GRPCServer *grpcapp.App
	}
)

func NewApp(
	log *slog.Logger,
)

func New(
	log *slog.Logger,
	grpcPort int64,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище (storage)

	// TODO: init auth service (auth)
	authService := auth.New(log, nil, nil, nil, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
