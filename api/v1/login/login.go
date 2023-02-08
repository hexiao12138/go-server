package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/model/common/response"
	"go-server/model/login/request"
	res "go-server/model/login/response"
)

type LoginApi struct{}

func (l *LoginApi) UsernameLogin(c *gin.Context) {
	var user request.LoginUser
	c.ShouldBindJSON(&user)
	err, userInfo := userService.UserNameLogin(user)
	if err != nil {
		response.FailWithData(res.UserLoignRes{}, "用户名或密码错误", c)
	} else {
		response.OkWithData(userInfo, "登录成功", c)
	}
}
func (l *LoginApi) MobileLogin(c *gin.Context) {
	var user request.LoginUser
	c.ShouldBindJSON(&user)
	fmt.Println(user)

}

// func (l *LoginApi) CreateJwtToken(user request.UserNameReq) {
// 	j := utils.JWT{}
// 	cliams := utils.MyClaims{
// 		Username: user.UserName,
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:    "hexiao",
// 			ExpiresAt: time.Now().Unix() + 60*60*60,
// 		},
// 	}
// 	token, err := j.CreateRsJWToken(cliams)
// 	if err != nil {
// 		fmt.Println("生成token失败", err)
// 		return
// 	}
// }
