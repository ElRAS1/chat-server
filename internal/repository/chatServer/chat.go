package chatServer

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

const (
	chatName      = "chats" //  Db name
	chatMessages  = "chat_messages"
	chatID        = "chat_id"
	chatUsernames = "usernames"    //  Store array usernames
	chatTimestamp = "timestamp"    // Time created chat
	chatReceiver  = "receiver"     // receiver message
	chatText      = "text_message" // message text
)

const (
	deadline = 5
)

type repo struct {
	Db     *pgxpool.Pool
	logger *slog.Logger
}

func New(dbClient *pgxpool.Pool, log *slog.Logger) *repo {
	return &repo{
		Db:     dbClient,
		logger: log,
	}
}
