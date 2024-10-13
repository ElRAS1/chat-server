package chatServer

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/chat-server/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (s *repo) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Insert(chatDbName).
		Columns(chatUsernames, chatCreatedAt, chatUpdatedAt).
		Values(req.Usernames, time.Now(), time.Now()).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("repo Create error: %v", err)
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo Create error: %v", err)
	}
	defer conn.Release()

	var id int64
	err = conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("repo Create error: %w", err)
	}

	return &model.CreateResponse{Id: id}, nil
}

func (s *repo) Delete(ctx context.Context, req *model.DeleteRequest) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Delete(chatDbName).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("repo : failed to build delete query: %w", err)
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("repo : failed to acquire database connection: %w", err)
	}
	defer conn.Release()

	res, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("repo : failed to execute delete query: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("repo : record not found")
	}

	return nil
}

func (s *repo) SendMessage(ctx context.Context, req *model.SendMessageRequest) error {

	return nil
}
