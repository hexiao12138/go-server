package initialize

import (
	"fmt"
	"go-server/middleware"
	"go-server/routers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 获取路由组实例
	loginRouter := routers.RouterGroupWeb.Login
	chatRouter := routers.RouterGroupWeb.Chat
	userRouter := routers.RouterGroupWeb.User
	PublicGroup := Router.Group("v1")
	PrivateGroup := Router.Group("").Use(middleware.JWTAuthMiddleware())
	fmt.Println(PrivateGroup)
	{
		loginRouter.InitLoginRouter(PublicGroup) // 注册登录路由
		chatRouter.InitChatRouter(PublicGroup)   // 注册chat路由
		userRouter.InitUserRouter(PublicGroup)   // 注册用户路由
	}
	return Router
}
