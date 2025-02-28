package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	"aiServer/core"
	"aiServer/global"
	"aiServer/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       AI-Server Swagger API接口文档
// @version                     v1.0.0
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.AI_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.AI_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.AI_LOG)
	global.AI_DB = initialize.Gorm() // gorm连接数据库
	if global.AI_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.AI_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
