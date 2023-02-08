package login

import (
	"fmt"
	v1 "go-server/api/v1"
	"go-server/global"
	"go-server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ResInfo struct {
	Code int      `json:"code"`
	Data LoginRes `json:"data"`
}
type LoginRes struct {
	Token string `json:"token"`
}

func Login(e *gin.Engine) {
	var user UserInfo
	e.POST("/v1/login", func(c *gin.Context) {
		j := utils.JWT{}
		cliams := utils.MyClaims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				Issuer:    "hexiao",
				ExpiresAt: time.Now().Unix() + 60*60*60,
			},
		}
		token, err := j.CreateRsJWToken(cliams)
		if err != nil {
			fmt.Println("生成token失败", err)
			return
		}
		err = c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println("err", err)
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(200, gin.H{
					"msg":  err.Error(),
					"code": 0,
				},
				)
			} else {
				c.JSON(200, gin.H{
					"msg":  errs.Translate(global.Trans),
					"code": 0,
				},
				)
			}
		} else {
			// global.Db.Create(&user)
			c.JSON(200, ResInfo{
				Code: 1,
				Data: LoginRes{
					Token: token,
				},
			})
		}
		fmt.Println("err", err)
	})
}

type LoginRouter struct{}

func (l *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	loginGroup := Router.Group("site")
	loginApi := v1.ApiGroupWeb.LoginApiGroup
	{
		loginGroup.POST("usernameLogin", loginApi.UsernameLogin)
		loginGroup.POST("mobileLogin", loginApi.UsernameLogin)
	}
}
