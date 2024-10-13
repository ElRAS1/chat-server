package chatServer

import "github.com/jackc/pgx/v5/pgxpool"

const (
	chatDbName    = "chats"      //  Db name
	chatUsernames = "usernames"  //  Store array usernames
	chatCreatedAt = "created_at" // Time created chat
	chatUpdatedAt = "updated_at" // Time updated chat

)

const (
	deadline = 5
)

type repo struct {
	Db *pgxpool.Pool
}

func New(dbClient *pgxpool.Pool) *repo {
	return &repo{
		Db: dbClient,
	}
}
