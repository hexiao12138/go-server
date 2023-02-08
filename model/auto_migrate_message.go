package model

import "time"

type MessageData struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	FromId   uint      `json:"fromId"`
	FromName string    `json:"fromName"`
	ToId     uint      `json:"toId"`
	ToName   string    `json:"toName"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
}
