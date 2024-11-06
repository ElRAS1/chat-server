package api

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ELRAS1/chat-server/internal/converter"
	"github.com/ELRAS1/chat-server/internal/service"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Api struct {
	*chatServer.UnimplementedChatServerServer
	serv   service.ChatServer
	logger slog.Logger
}

func New(srv service.ChatServer, log *slog.Logger) *Api {
	return &Api{
		serv:                          srv,
		UnimplementedChatServerServer: &chatServer.UnimplementedChatServerServer{},
		logger:                        *log,
	}
}

func (a *Api) Create(ctx context.Context, req *chatServer.CreateRequest) (*chatServer.CreateResponse, error) {
	if len(req.GetUsernames()) < 2 {
		err := errors.New("the number of participants in the chat must be at least 2")
		a.logger.Debug(err.Error())

		return nil, err
	}

	resp, err := a.serv.Create(ctx, converter.ApiCreateToModel(req))
	if err != nil {
		a.logger.Debug(err.Error())

		return nil, err
	}

	return converter.ModelCreateToApi(resp), nil
}

func (a *Api) Delete(ctx context.Context, req *chatServer.DeleteRequest) (*emptypb.Empty, error) {
	if err := a.serv.Delete(ctx, converter.ApiDeleteToModel(req)); err != nil {
		a.logger.Debug(err.Error())

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (a *Api) SendMessage(ctx context.Context, req *chatServer.SendMessageRequest) (*emptypb.Empty, error) {
	if req.GetFrom() == "" {
		err := errors.New("'From' field cannot be empty")
		a.logger.Debug(err.Error())

		return nil, err
	}

	if req.GetText() == "" {
		err := errors.New("'Text' field cannot be empty")
		a.logger.Debug(err.Error())

		return nil, err
	}

	if err := a.serv.SendMessage(ctx, converter.ApiSendMessageToModel(req)); err != nil {
		a.logger.Debug(err.Error())

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
