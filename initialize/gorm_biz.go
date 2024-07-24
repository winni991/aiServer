package initialize

import (
	"aiServer/global"
)

func bizModel() error {
	db := global.AI_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
