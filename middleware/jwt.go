package middleware

import (
	"fmt"
	"go-server/utils"

	"github.com/gin-gonic/gin"
)

// 校验jwt中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(200, gin.H{
				"msg":  "未登录或非法访问",
				"code": 401,
			})
			c.Abort()
			return
		}
		j := utils.JWT{}
		claims, err := j.ParseRsJWToken(c.Request.Header.Get("x-token"), &utils.MyClaims{})
		if err != nil {
			fmt.Println("出错了", err)
			c.JSON(200, gin.H{
				"msg":  "您的账户异地登录或令牌失效，请重新登录",
				"code": 401,
			})
			c.Abort()
			return
		}
		fmt.Println(12, claims, err)
	}
}
