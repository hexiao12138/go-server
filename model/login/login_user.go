package login

import (
	"gorm.io/gorm"
	"time"
)

type LoginUser struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserName  string         `json:"userName"`
	Mobile    string         `json:"mobile"`
}
