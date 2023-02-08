package user

import (
	"go-server/global"
	"go-server/model/chat/request"
	"go-server/model/login"
)

type User struct {
}

func (u *User) GetUserInfo(userInfo request.SearchFriend) (error, *login.LoginUser) {
	var user login.LoginUser
	err := global.My_DB.Where("user_name = ? or mobile = ?", userInfo.UserName, userInfo.UserName).First(&user).Error
	return err, &user
}
