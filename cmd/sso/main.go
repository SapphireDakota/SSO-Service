package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
)

func main() {
	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.Any("config", *cfg),
	)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
