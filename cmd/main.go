package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/chat-server/internal/config"
	"github.com/ELRAS1/chat-server/internal/server"
	"github.com/ELRAS1/chat-server/pkg/logger"
)

func main() {
	cfg, err := config.NewServerCfg()
	if err != nil {
		log.Fatalln(err)
	}

	logger := logger.ConfigureLogger(cfg.LogLevel, cfg.ConfigLog)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server, listener, storage, err := server.New(ctx, logger, *cfg)
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info(fmt.Sprintf("app starting in port %s", cfg.Port))
	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	storage.Close()
	logger.Info("app closed...")
}
