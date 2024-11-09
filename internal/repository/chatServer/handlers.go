package chatServer

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ELRAS1/chat-server/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (s *repo) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Insert(chatName).
		Columns(chatUsernames).
		Values(req.Usernames).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		err = fmt.Errorf("repo Create error: %w", err)
		s.logger.Debug(err.Error())

		return nil, err
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("repo Create error: %w", err)
		s.logger.Debug(err.Error())

		return nil, err
	}
	defer conn.Release()

	var id int64
	err = conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		err = fmt.Errorf("repo Create error: %w", err)
		s.logger.Debug(err.Error())

		return nil, err
	}

	s.logger.Info("chat was created successfully: [participants]: ",
		strings.Join(req.Usernames, ", "))

	return &model.CreateResponse{Id: id}, nil
}

func (s *repo) Delete(ctx context.Context, req *model.DeleteRequest) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	query, args, err := sq.Delete(chatName).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("repo : failed to build delete query: %w", err)
		s.logger.Debug(err.Error())

		return err
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("repo : failed to acquire database connection: %w", err)
		s.logger.Debug(err.Error())

		return err
	}
	defer conn.Release()

	res, err := conn.Exec(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("repo : failed to execute delete query: %w", err)
		s.logger.Debug(err.Error())

		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		err = errors.New("repo : record not found")
		s.logger.Debug(err.Error())

		return err
	}

	s.logger.Info("chat was deleted successfully", req.Id)

	return nil
}

func (s *repo) SendMessage(ctx context.Context, req *model.SendMessageRequest) error {
	query, args, err := sq.Insert(chatMessages).
		Columns(chatID, chatReceiver, chatText, chatTimestamp).
		Values(req.ChatId, req.From, req.Text, req.Timestamp).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("repo SendMessage error: %w", err)
		s.logger.Debug(err.Error())

		return err
	}

	conn, err := s.Db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("repo : failed to acquire database connection: %w", err)
		s.logger.Debug(err.Error())

		return err
	}
	defer conn.Release()

	if _, err = conn.Exec(ctx, query, args...); err != nil {
		err = fmt.Errorf("repo SendMessage error: %v", err)
		s.logger.Debug(err.Error())

		return err
	}

	s.logger.Info("chat was sent successfully to ", req.From)

	return nil
}
