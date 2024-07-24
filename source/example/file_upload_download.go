package example

import (
	"aiServer/model/example"
	"aiServer/service/system"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderExaFile = system.InitOrderInternal + 1

type initExaFileMysql struct{}

// auto run
func init() {
	system.RegisterInit(initOrderExaFile, &initExaFileMysql{})
}

func (i *initExaFileMysql) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&example.ExaFileUploadAndDownload{})
}

func (i *initExaFileMysql) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&example.ExaFileUploadAndDownload{})
}

func (i initExaFileMysql) InitializerName() string {
	return example.ExaFileUploadAndDownload{}.TableName()
}

func (i *initExaFileMysql) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []example.ExaFileUploadAndDownload{
		{Name: "girl.jpeg", Url: "uploads/file/28ca53d2b7bb4aa13549b4022c79dca1_20240723191403.jpeg", Tag: "jpeg", Key: "28ca53d2b7bb4aa13549b4022c79dca1_20240723191403.jpeg"},
		{Name: "boy.jpeg", Url: "uploads/file/1a699ad5e06aa8a6db3bcf9cfb2f00f2_20240723191338.jpeg", Tag: "jpeg", Key: "1a699ad5e06aa8a6db3bcf9cfb2f00f2_20240723191338.jpeg"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, example.ExaFileUploadAndDownload{}.TableName()+"表数据初始化失败!")
	}
	return ctx, nil
}

func (i *initExaFileMysql) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	lookup := example.ExaFileUploadAndDownload{Name: "boy.jpeg", Key: "1a699ad5e06aa8a6db3bcf9cfb2f00f2_20240723191338.jpeg"}
	if errors.Is(db.First(&lookup, &lookup).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
