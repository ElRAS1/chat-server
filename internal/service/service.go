package service

import (
	"context"
	"github.com/ELRAS1/chat-server/internal/model"
)

type ChatServer interface {
	Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error)
	Delete(ctx context.Context, req *model.DeleteRequest) error
	SendMessage(ctx context.Context, req *model.SendMessageRequest) error
}
