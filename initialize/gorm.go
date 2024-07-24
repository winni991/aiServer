package initialize

import (
	"os"

	"aiServer/global"
	"aiServer/model/example"
	"aiServer/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables() {
	db := global.AI_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysOperationRecord{},

		example.ExaFile{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.AI_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.AI_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.AI_LOG.Info("register table success")
}
