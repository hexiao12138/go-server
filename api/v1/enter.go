package v1

import (
	"go-server/api/v1/chat"
	"go-server/api/v1/login"
	"go-server/api/v1/user"
)

type ApiGroup struct {
	LoginApiGroup login.LoginApi
	ChatApiGroup  chat.ChatApi
	UserApiGroup  user.UserApi
}

var ApiGroupWeb = new(ApiGroup)
