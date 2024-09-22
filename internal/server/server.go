package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/ELRAS1/chat-server/internal/config"
	"github.com/ELRAS1/chat-server/internal/handlers"
	"github.com/ELRAS1/chat-server/internal/storage"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(ctx context.Context, logger *slog.Logger, cfg config.Server) (*grpc.Server, net.Listener, *storage.Storage, error) {
	lis, err := net.Listen(cfg.Network, cfg.Port)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w", err)
	}

	server := handlers.New(logger)

	logger.Info("Starting database connection...")
	server.Db, err = storage.ConfigureStorage(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w", err)
	}
	logger.Info("Successful connection to the db")

	srv := grpc.NewServer()
	reflection.Register(srv)
	chatServer.RegisterChatServerServer(srv, server)

	return srv, lis, server.Db, nil
}
