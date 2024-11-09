package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
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
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	logger := logger.New(cfg.LogLevel, cfg.ConfigLog)
	slog.SetDefault(logger)

	listener, err := net.Listen(cfg.Network, cfg.GRPCPort)
	if err != nil {
		log.Fatalln(err)
	}

	dbClient, err := config.InitializeDatabaseClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	repository := repo.New(dbClient, logger)

	service := serv.New(repository)

	server := grpc.NewServer()
	reflection.Register(server)
	chatServer.RegisterChatServerServer(server, api.New(service, logger))

	go func() {
		logger.Info(fmt.Sprintf("grpc server is running on: %v", net.JoinHostPort(cfg.Host, cfg.GRPCPort)))
		if err = server.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	httpServer := config.InitHTTP(ctx, cfg.GRPCPort, cfg.HTTPPort)
	go func() {
		logger.Info(fmt.Sprintf("http server is running on: %v", net.JoinHostPort(cfg.Host, cfg.HTTPPort)))
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalln(fmt.Sprintf("failed to http serve: %v", err))
		}
	}()

	httpSwagger := config.InitSwagger()
	go func() {
		logger.Info(fmt.Sprintf("swagger ui is running on: %v", net.JoinHostPort(cfg.Host, cfg.SwaggerPort)))
		if err = http.ListenAndServe(cfg.SwaggerPort, httpSwagger); err != nil {
			log.Fatalln(fmt.Sprintf("failed to swagger serve: %v", err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
