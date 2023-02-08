package core

import (
	"fmt"
	"go-server/initialize"
)

func RunServer() {
	Router := initialize.Routers()
	fmt.Println("服务启动成功")
	Router.Run()
}
