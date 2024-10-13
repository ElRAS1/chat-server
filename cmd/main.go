package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/chat-server/internal/api"
	"github.com/ELRAS1/chat-server/internal/config"
	repo "github.com/ELRAS1/chat-server/internal/repository/chatServer"
	serv "github.com/ELRAS1/chat-server/internal/service/chatServer"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"github.com/ELRAS1/chat-server/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.NewServerCfg()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.New(cfg.LogLevel, cfg.ConfigLog)
	slog.SetDefault(logger)

	listener, err := net.Listen(cfg.Network, cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	dbClient, err := config.InitializeDatabaseClient(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	repository := repo.New(dbClient)

	service := serv.New(repository)

	server := grpc.NewServer()
	reflection.Register(server)
	chatServer.RegisterChatServerServer(server, api.New(service))

	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	logger.Info(fmt.Sprintf("the server is running on the port %v", cfg.Port))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
