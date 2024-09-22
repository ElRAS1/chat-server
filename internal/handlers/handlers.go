package handlers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ELRAS1/chat-server/internal/storage"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	*chatServer.UnimplementedChatServerServer
	Db     *storage.Storage
	Logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		Logger:                        logger,
		UnimplementedChatServerServer: &chatServer.UnimplementedChatServerServer{},
	}
}

func (h *Handler) Create(ctx context.Context, req *chatServer.CreateRequest) (*chatServer.CreateResponse, error) {
	if len(req.Usernames) < 2 {
		err := fmt.Errorf("the number of participants in the chat must be at least 2")
		h.Logger.Error(err.Error())
		return nil, err
	}

	id, err := h.Db.CreateChat(ctx, req)
	if err != nil {
		h.Logger.Error(err.Error())
		return nil, err
	}

	return &chatServer.CreateResponse{Id: id}, nil
}

func (h *Handler) Delete(ctx context.Context, req *chatServer.DeleteRequest) (*empty.Empty, error) {
	_, err := h.Db.DeleteChat(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}
func (h *Handler) SendMessage(ctx context.Context, req *chatServer.SendMessageRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}
