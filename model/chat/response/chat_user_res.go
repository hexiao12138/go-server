package response

import (
	"go-server/model"
	"go-server/model/login"
)

type SearchFriendRes struct {
	UserInfo *login.LoginUser `json:"userInfo"`
}

type ChatMessage struct {
	List      []model.MessageData `json:"list"`
	MessageId uint                `json:"messageId"`
}
