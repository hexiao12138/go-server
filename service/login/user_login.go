package login

import (
	"go-server/global"
	"go-server/model"
	res "go-server/model/chat/response"
	"go-server/model/login"
	"go-server/model/login/request"
	"go-server/model/login/response"
	"go-server/utils"
	"sort"
)

type UserLogin struct {
}

func (u *UserLogin) UserNameLogin(userInfo request.LoginUser) (error, *response.UserLoignRes) {
	var user login.LoginUser
	var friends []model.UserFriends
	var users []*login.LoginUser
	var friendInfo *login.LoginUser
	var messageList []model.MessageData
	var messageResList []res.ChatMessage
	err := global.My_DB.Where("user_name = ? and pass_word = ?", userInfo.UserName, userInfo.PassWord).First(&user).Error
	if err == nil {
		global.My_DB.Where("from_id = ?", user.ID).Find(&messageList)
		sort.Slice(messageList, func(i, j int) bool { return messageList[i].ToId < messageList[j].ToId })
		messageResList = utils.GroupMessage(messageList, "my")
		global.My_DB.Where("user_id = ? ", user.ID).Find(&friends)
		for _, friend := range friends {
			global.My_DB.Where("id = ? ", friend.FriendId).Find(&friendInfo)
			users = append(users, friendInfo)
			friendInfo = nil
		}
	}
	return err, &response.UserLoignRes{
		Token:      "333",
		UserInfo:   &user,
		Friends:    users,
		MessageMap: messageResList,
	}
}
