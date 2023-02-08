package service

import (
	"go-server/service/chat"
	"go-server/service/login"
)

type ServiceGroup struct {
	ServiceLoginGroup login.UserLogin
	ServiceChatGroup  chat.Chat
}

var ServiceGroupWeb = new(ServiceGroup)
