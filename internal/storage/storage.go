package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/chat-server/pkg/chatServer"
	sq "github.com/Masterminds/squirrel"
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

	err = s.Db.QueryRow(ctx, query, args...).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("CreateChat error: %v", err)
	}

	return
}
