package global

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"aiServer/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	AI_DB     *gorm.DB
	AI_REDIS  redis.UniversalClient
	AI_MONGO  *qmgo.QmgoClient
	AI_CONFIG config.Server
	AI_VP     *viper.Viper
	// AI_LOG    *oplogging.Logger
	AI_LOG                 *zap.Logger
	AI_Concurrency_Control = &singleflight.Group{}
	AI_ROUTERS             gin.RoutesInfo
	BlackCache             local_cache.Cache
	lock                   sync.RWMutex
)
