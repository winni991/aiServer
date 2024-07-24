package initialize

import (
	"context"

	"aiServer/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.AI_CONFIG.Redis
	var client redis.UniversalClient
	// 使用单例模式
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.AI_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.AI_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.AI_REDIS = client
	}
}
