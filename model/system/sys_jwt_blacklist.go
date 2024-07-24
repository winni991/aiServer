package system

import (
	"aiServer/global"
)

type JwtBlacklist struct {
	global.AI_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
