package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/ELRAS1/chat-server/internal/config"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db  *pgxpool.Pool
	Cfg *config.Db
}

// ConfigureStorage: configurations and database connection
func ConfigureStorage(ctx context.Context) (*Storage, error) {
	const nm = "[ConfigureStorage]"

	connStr, err := newConnStr()
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	cfg, err := config.NewDbCfg()
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	if err = db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	return &Storage{Db: db, Cfg: cfg}, nil
}

func newConnStr() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("no .env file found")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	return connStr, nil
}

func (s *Storage) Close() {
	s.Db.Close()
}
