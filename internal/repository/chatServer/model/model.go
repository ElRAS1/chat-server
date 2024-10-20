package model

import "time"

type CreateRequest struct {
	Usernames []string `db:"usernames"`
}
type CreateResponse struct {
	Id int64 `db:"id"`
}

type DeleteRequest struct {
	Id int64 `db:"id"`
}

type SendMessageRequest struct {
	ChatId    int64     `db:"chat_id"`
	From      string    `db:"from"`
	Text      string    `db:"text"`
	Timestamp time.Time `db:"timestamp"`
}
