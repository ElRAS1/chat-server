package logger

import (
	"log/slog"
	"os"
)

func New(level int, cfg string) (logger *slog.Logger) {
	switch cfg {
	case "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	}

	return logger
}
