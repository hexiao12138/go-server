package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Phone     string
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
}
type Client struct {
	wsChan chan interface{}
}
type UserFriends struct {
	ID       uint `json:"id" gorm:"primarykey"`
	UserId   int
	FriendId int
}
