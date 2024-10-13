package model

import "time"

type CreateRequest struct {
	Usernames []string
}
type CreateResponse struct {
	Id int64
}

type DeleteRequest struct {
	Id int64
}

type SendMessageRequest struct {
	From      string
	Text      string
	Timestamp time.Time
}