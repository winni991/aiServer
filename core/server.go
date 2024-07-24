package core

import (
	"aiServer/global"
	"aiServer/initialize"
	"aiServer/service/system"
	"fmt"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.AI_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	// 从db加载jwt数据
	if global.AI_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.AI_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.AI_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.AI_LOG.Error(s.ListenAndServe().Error())
}
