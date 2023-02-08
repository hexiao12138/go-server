package user

import (
	"github.com/gin-gonic/gin"
	"go-server/model/chat/request"
	res "go-server/model/chat/response"
	"go-server/model/common/response"
)

type UserApi struct {
}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	var user request.SearchFriend
	c.ShouldBindQuery(&user)
	err, userInfo := userService.GetUserInfo(user)
	if err != nil {
		response.FailWithData(res.SearchFriendRes{}, "该用户不存在", c)
	} else {
		response.OkWithData(res.SearchFriendRes{UserInfo: userInfo}, "", c)
	}
}
