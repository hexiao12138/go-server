package chat

import (
	v1 "go-server/api/v1"
	"go-server/api/v1/chat"

	"github.com/gin-gonic/gin"
)

type ChatWsRouter struct {
}

func (w *ChatWsRouter) InitChatRouter(Router *gin.RouterGroup) {
	hub := chat.NewHub()
	go hub.Run()
	chatGroup := Router.Group("chat")
	chatApi := v1.ApiGroupWeb.ChatApiGroup
	{
		chatGroup.POST("friends", chatApi.AddFriend)
		chatGroup.GET("friends", chatApi.SearchFriend)
		chatGroup.GET("connWs", chatApi.ServerWsHander(hub))
	}
}
