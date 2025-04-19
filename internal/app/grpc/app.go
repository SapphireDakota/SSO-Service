package grpcapp

import (
	"log/slog"
	"net"
	authgrpc "sso/internal/auth"
	"strconv"

	"google.golang.org/grpc"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int64
}

func New(log *slog.Logger, port int64) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		log:  log,
		gRPC: gRPCServer,
		port: port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	a.log.Info("starting grpc server")

	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(a.port, 10))
	if err != nil {
		return err
	}

	a.log.Info("grpc server is listening", slog.Int64("port", a.port))

	if err := a.gRPC.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	a.log.Info("stopping grpc server")

	a.gRPC.GracefulStop()
}
