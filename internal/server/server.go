package server

import (
	"log/slog"
	"net"

	"github.com/ELRAS1/chat-server/internal/config"
	"github.com/ELRAS1/chat-server/internal/handlers"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(logger *slog.Logger, cfg config.Config) (srv *grpc.Server, lis net.Listener, err error) {
	lis, err = net.Listen(cfg.Network, cfg.Port)
	if err != nil {
		return nil, nil, err
	}

	server := handlers.New(logger)

	srv = grpc.NewServer()
	reflection.Register(srv)
	chatServer.RegisterChatServerServer(srv, server)

	return
}
