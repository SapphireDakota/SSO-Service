package app

import (
	"google.golang.org/grpc"
	"log/slog"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port string
}
