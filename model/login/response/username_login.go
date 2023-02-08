package response

import (
	"go-server/model"
	"go-server/model/chat/response"
	"go-server/model/login"
)

type UserLoignRes struct {
	Token       string                 `json:"token"`
	UserInfo    *login.LoginUser       `json:"userInfo"`
	Friends     []*login.LoginUser     `json:"friends"`
	MessageList []model.MessageData    `json:"messageList"`
	MessageMap  []response.ChatMessage `json:"messageListData"`
}
