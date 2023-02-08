package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id int `json:"id"`
}
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 0
	SUCCESS = 1
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OkWithData(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}
func FailWithData(data interface{}, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}
