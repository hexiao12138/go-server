package chat

import (
	"go-server/model"
	"go-server/model/chat/request"
	res "go-server/model/chat/response"
	"go-server/model/common/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatApi struct {
}

// func (ch *ChatApi) ServerWs(c *gin.Context) {
// 	// 升级ws协议
// 	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Println("连接失败")
// 		return
// 	}
// 	fmt.Println(wsConn)
// }
func (ch *ChatApi) SearchFriend(c *gin.Context) {
	var user request.SearchFriend
	c.ShouldBindQuery(&user)
	err, userInfo := chatService.SearchFriend(user)
	if err != nil {
		response.FailWithData(res.SearchFriendRes{}, "该用户不存在", c)
	} else {
		response.OkWithData(res.SearchFriendRes{UserInfo: userInfo}, "", c)
	}
}
func (ch *ChatApi) AddFriend(c *gin.Context) {
	var data model.UserFriends
	c.ShouldBindJSON(&data)
	err := chatService.AddFriend(&data)
	if err != nil {
		response.FailWithData(res.SearchFriendRes{}, "添加用户失败", c)
	} else {
		response.OkWithData(res.SearchFriendRes{}, "", c)
	}
}
func (ch *ChatApi) ServerWsHander(hub *Hub) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{
			hub:  hub,
			conn: conn,
			send: make(chan []byte, bufSize),
		}
		client.hub.register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		go client.readPump()
	}
	return gin.HandlerFunc(fn)
}
func (ch *ChatApi) ServeWs(c *gin.Context) {
	hub := NewHub()
	go hub.Run()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, bufSize),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
