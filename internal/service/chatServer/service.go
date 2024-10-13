package chatServer

import (
	"context"
	"github.com/ELRAS1/chat-server/internal/model"
	"github.com/ELRAS1/chat-server/internal/repository"
)

type service struct {
	chatServerRepo repository.RepoChatServer
}

func New(repo repository.RepoChatServer) *service {
	return &service{
		chatServerRepo: repo,
	}
}

func (s *service) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	return s.chatServerRepo.Create(ctx, req)
}

func (s *service) Delete(ctx context.Context, req *model.DeleteRequest) error {
	return s.chatServerRepo.Delete(ctx, req)
}

func (s *service) SendMessage(ctx context.Context, req *model.SendMessageRequest) error {
	return s.chatServerRepo.SendMessage(ctx, req)
}
