package api

import (
	"context"
	"fmt"

	"github.com/ELRAS1/chat-server/internal/converter"
	"github.com/ELRAS1/chat-server/internal/service"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Api struct {
	*chatServer.UnimplementedChatServerServer
	serv service.ChatServer
}

func New(srv service.ChatServer) *Api {
	return &Api{
		serv:                          srv,
		UnimplementedChatServerServer: &chatServer.UnimplementedChatServerServer{},
	}
}

func (a *Api) Create(ctx context.Context, req *chatServer.CreateRequest) (*chatServer.CreateResponse, error) {
	if len(req.Usernames) < 2 {
		err := fmt.Errorf("the number of participants in the chat must be at least 2")
		return nil, err
	}

	resp, err := a.serv.Create(ctx, converter.ApiCreateToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.ModelCreateToApi(resp), nil
}

func (a *Api) Delete(ctx context.Context, req *chatServer.DeleteRequest) (*emptypb.Empty, error) {
	if err := a.serv.Delete(ctx, converter.ApiDeleteToModel(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (a *Api) SendMessage(ctx context.Context, req *chatServer.SendMessageRequest) (*emptypb.Empty, error) {
	if req.From == "" {
		return nil, fmt.Errorf("'From' field cannot be empty")
	}

	if req.Text == "" {
		return nil, fmt.Errorf("'Text' field cannot be empty")
	}

	if err := a.serv.SendMessage(ctx, converter.ApiSendMessageToModel(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
