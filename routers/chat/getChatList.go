package chat

import (
	"encoding/json"
	"fmt"
	"go-server/middleware"
	"net/http"
	"unsafe"

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
var ReadChan chan WsUser

type WsUser struct {
	MyName  string `json:"myName"`
	NewName string `json:"newName"`
}

func GetChatList(e *gin.Engine) {
	e.GET("/", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":    "2323",
			"status": 0,
		})
	})
}
func ConnectWs(e *gin.Engine) {
	e.GET("/v1/wsChat", func(c *gin.Context) {
		// 升级ws协议
		ReadChan = make(chan WsUser, 1000)
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		fmt.Println("连接成功")
		go ReadClientMessage(wsConn)
		go WriteClientMessage(wsConn)
		if err != nil {
			fmt.Println(333, err)
			panic("连接出错退出服务")
		}

	})
}
func ReadClientMessage(wsConn *websocket.Conn) {
	for {
		_, message, wsErr := wsConn.ReadMessage()
		if wsErr != nil {
			fmt.Println("读取失败", wsErr)
			wsConn.Close()
			return
			// panic("读取失败")
		}
		wsUser := WsUser{}
		jErr := json.Unmarshal(message, &wsUser)
		if jErr != nil {
			fmt.Println("失败")
		}
		ReadChan <- wsUser
	}
}
func WriteClientMessage(wsConn *websocket.Conn) {
	for {
		select {
		case data := <-ReadChan:
			fmt.Println(12, data.MyName)
			json, err := json.Marshal(data)
			if err != nil {
				fmt.Println("出错了", err)
			}
			wsConn.WriteMessage(1, json)
		}
	}
}
func byteSliceToString(bytes []byte) string {

	return *(*string)(unsafe.Pointer(&bytes))

}
