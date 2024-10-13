package converter

import (
	"github.com/ELRAS1/chat-server/internal/model"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
)

func ModelCreateToApi(req *model.CreateResponse) *chatServer.CreateResponse {
	return &chatServer.CreateResponse{
		Id: req.Id,
	}
}
