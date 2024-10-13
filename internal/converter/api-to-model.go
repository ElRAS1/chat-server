package converter

import (
	"github.com/ELRAS1/chat-server/internal/model"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
)

func ApiCreateToModel(req *chatServer.CreateRequest) *model.CreateRequest {
	return &model.CreateRequest{
		Usernames: req.Usernames,
	}
}

func ApiDeleteToModel(req *chatServer.DeleteRequest) *model.DeleteRequest {
	return &model.DeleteRequest{
		Id: req.Id,
	}
}

func ApiSendMessageToModel(req *chatServer.SendMessageRequest) *model.SendMessageRequest {
	return &model.SendMessageRequest{
		From:      req.From,
		Text:      req.Text,
		Timestamp: req.Timestamp.AsTime(),
	}
}
