package user

import (
	"github.com/gin-gonic/gin"
	v1 "go-server/api/v1"
)

type UserRouter struct{}

func (l *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("user")
	userApi := v1.ApiGroupWeb.UserApiGroup
	{
		UserGroup.GET("userInfo", userApi.GetUserInfo)
	}
}
