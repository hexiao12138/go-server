package utils

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type JWT struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// 根据rs256生成token
func (j *JWT) CreateRsJWToken(claims MyClaims) (string, error) {
	path, _ := os.Getwd()
	privateKeyByte, err := os.ReadFile(path + "/private.key")
	if err != nil {
		fmt.Println("读取失败", err)
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	j.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		fmt.Println(9999, err)
	}
	return t.SignedString(j.privateKey)
}

// 根据rs256解析token
func (j *JWT) ParseRsJWToken(tokenString string, claims *MyClaims) (*MyClaims, error) {
	path, _ := os.Getwd()
	publicKeyByte, _ := os.ReadFile(path + "/public.key")
	j.publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	} else {
		return token.Claims.(*MyClaims), nil
	}
}
