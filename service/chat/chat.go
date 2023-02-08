package chat

import (
	"go-server/global"
	"go-server/model"
	"go-server/model/chat/request"
	"go-server/model/login"
)

type Chat struct {
}

func (c *Chat) AddFriend(userInfo *model.UserFriends) error {
	err := global.My_DB.Create(userInfo).Error
	return err
}
func (c *Chat) SearchFriend(userInfo request.SearchFriend) (error, *login.LoginUser) {
	var user login.LoginUser
	subQuery := global.My_DB.Model(&login.LoginUser{})
	err := global.My_DB.Table("login_users").Where("id = ?", subQuery).Find(&user).Error
	return err, &user
}
