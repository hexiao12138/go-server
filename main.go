package main

import (
	"crypto/rsa"
	"go-server/core"
	"go-server/global"
	"go-server/initialize"

	"github.com/golang-jwt/jwt"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	global.My_DB = initialize.GormMysql()
	if global.My_DB != nil {
		initialize.RegisterTables(global.My_DB)
	}
	core.RunServer()
}
