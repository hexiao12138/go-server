package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "io"
	"os"
	"time"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

type UserInfo struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
}
type ResInfo struct {
	Data LoginRes `json:"data"`
}
type LoginRes struct {
	Token string `json:"token"`
}

func main() {
	var user UserInfo
	j := JWT{
		SigningKey: []byte("hellogolang"),
	}
	dsn := "root:8524xiao@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate((&UserInfo{}))
	r := gin.Default()
	r.POST("/api/login", func(c *gin.Context) {
		c.ShouldBindJSON(&user)
		claims := MyClaims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				Issuer:    "hexiao",
				ExpiresAt: time.Now().Unix() + 60*60*60,
			},
		}
		token, err := CreateRsJWToken(claims)
		// db.Create(&user)
		if err == nil {
			c.JSON(200, ResInfo{
				Data: LoginRes{
					Token: token,
				},
			})
		}
	})
	r.GET("api/getChatList", JWTAuthMiddleware(&j), func(c *gin.Context) {

	})
	r.Run()
}
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte("helloxJWT"),
	}
}

// 创建token
// func (j *JWT) CreateJWToken(claims MyClaims) (string, error) {
// 	fmt.Println("创建token")
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(j.SigningKey)
// }

// // 解析token
// func (j *JWT) ParseJWToken(tokenString string, claims *MyClaims) {
// 	fmt.Println("解析token")
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		fmt.Println(2222, err)
// 		return
// 	}
// 	fmt.Println(1111, token.Claims.(*MyClaims).Username)
// }

// 校验jwt中间件
func JWTAuthMiddleware(j *JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		ParseRsJWToken(c.Request.Header.Get("x-token"), &MyClaims{})
	}
}

// 根据rs256生成token
func CreateRsJWToken(claims MyClaims) (string, error) {
	privateKeyByte, err := os.ReadFile("./private.key")
	if err != nil {
		fmt.Println("读取失败", err)
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		fmt.Println(9999, err)
	}
	return t.SignedString(privateKey)
}

// 根据rs256解析token
func ParseRsJWToken(tokenString string, claims *MyClaims) {
	publicKeyByte, _ := os.ReadFile("./public.key")
	publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	fmt.Println("成功", publicKey)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		fmt.Println("token校验失败", err)
	} else {
		fmt.Println("token校验成功", token)
	}
}
