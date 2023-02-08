package service

import (
	"go-server/global"
	"go-server/model"
)

// 初始化表
func initTables() error {
	return global.My_DB.AutoMigrate(
		global.My_DB.AutoMigrate(&model.UserInfo{}),
		global.My_DB.AutoMigrate(&model.UserFriends{}),
	)
}
