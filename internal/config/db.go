package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	deadline = 5
)

func InitializeDatabaseClient(ctx context.Context) (*pgxpool.Pool, error) {
	const nm = "[InitializeDatabaseClient]"

	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	connStr, err := newConnStr()
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	if err = db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	return db, nil
}

func newConnStr() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("no .env file found")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	), nil
}
