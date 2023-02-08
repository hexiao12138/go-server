package initialize

import (
	"fmt"
	"go-server/model"
	"go-server/model/login"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3307)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接错误")
		return nil
	}
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&login.LoginUser{},
		&model.UserFriends{},
		&model.MessageData{},
	)
	if err != nil {
		fmt.Println("初始化表失败")
	}
}
