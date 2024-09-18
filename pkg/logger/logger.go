package logger

import (
	"log/slog"
	"os"
)

// ConfigureLogger настраивает логгер с предоставленной конфигурацией(configures the logger with the provided configuration).
func ConfigureLogger(level int, cfg string) (logger *slog.Logger) {
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
