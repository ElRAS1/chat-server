package handlers

import (
	"context"
	"log/slog"

	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	*chatServer.UnimplementedChatServerServer
	Logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		Logger:                        logger,
		UnimplementedChatServerServer: &chatServer.UnimplementedChatServerServer{},
	}
}

func (h *Handler) Create(ctx context.Context, req *chatServer.CreateRequest) (*chatServer.CreateResponse, error) {
	return &chatServer.CreateResponse{Id: 0}, nil
}
func (h *Handler) Delete(ctx context.Context, req *chatServer.DeleteRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (h *Handler) SendMessage(ctx context.Context, req *chatServer.SendMessageRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}
