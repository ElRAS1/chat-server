package server

import (
	"net"

	"github.com/ELRAS1/chat-server/internal/config"
	chatServer "github.com/ELRAS1/chat-server/pkg/chatServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	chatServer.UnimplementedChatServerServer
}

func New(cfg config.Config) (srv *grpc.Server, lis net.Listener, err error) {

	lis, err = net.Listen(cfg.Network, cfg.Port)
	if err != nil {
		return nil, nil, err
	}

	srv = grpc.NewServer()
	reflection.Register(srv)
	chatServer.RegisterChatServerServer(srv, server{})

	return
}
