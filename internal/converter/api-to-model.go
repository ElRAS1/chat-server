package converter

import (
	"github.com/ELRAS1/chat-server/internal/model"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
)

func ApiCreateToModel(req *chatServer.CreateRequest) *model.CreateRequest {
	return &model.CreateRequest{
		Usernames: req.GetUsernames(),
	}
}

func ApiDeleteToModel(req *chatServer.DeleteRequest) *model.DeleteRequest {
	return &model.DeleteRequest{
		Id: req.GetId(),
	}
}

func ApiSendMessageToModel(req *chatServer.SendMessageRequest) *model.SendMessageRequest {
	return &model.SendMessageRequest{
		ChatId:    req.GetChatId(),
		From:      req.GetFrom(),
		Text:      req.GetText(),
		Timestamp: req.Timestamp.AsTime(),
	}
}
