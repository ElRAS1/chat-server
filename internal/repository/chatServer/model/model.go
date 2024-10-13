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
	From      string    `db:"from"`
	Text      string    `db:"text"`
	Timestamp time.Time `db:"timestamp"`
}
