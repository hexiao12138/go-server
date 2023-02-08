package routers

import (
	"go-server/routers/chat"
	"go-server/routers/login"
	"go-server/routers/user"
)

type RouterGroup struct {
	Login login.LoginRouter
	Chat  chat.ChatWsRouter
	User  user.UserRouter
}

var RouterGroupWeb = new(RouterGroup)
