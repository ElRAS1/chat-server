package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/chat-server/pkg/chatServer"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	deadline = 5
)

func (s *Storage) CreateChat(ctx context.Context, req *chatServer.CreateRequest) (id int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Insert(s.Cfg.ChatDbName).
		Columns(s.Cfg.ChatUsernames, s.Cfg.ChatCreatedAt, s.Cfg.ChatUpdatedAt).
		Values(req.Usernames, time.Now(), time.Now()).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return 0, fmt.Errorf("CreateChat error: %v", err)
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("CreateChat error: %v", err)
	}

	return
}

func (s *Storage) DeleteChat(ctx context.Context, req *chatServer.DeleteRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Delete(s.Cfg.ChatDbName).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build delete query: %w", err)
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection: %w", err)
	}
	defer conn.Release()

	res, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return &emptypb.Empty{}, nil
}
