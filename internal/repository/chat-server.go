package repository

import (
	"context"

	"github.com/ELRAS1/chat-server/internal/model"
)

type RepoChatServer interface {
	Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error)
	Delete(ctx context.Context, req *model.DeleteRequest) error
	SendMessage(ctx context.Context, req *model.SendMessageRequest) error
}
