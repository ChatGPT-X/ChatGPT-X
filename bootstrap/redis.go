package bootstrap

import (
	"chatgpt_x/pkg/config"
	rds "chatgpt_x/pkg/redis"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// SetupRedis 初始化 Redis.
func SetupRedis() {
	addr := fmt.Sprintf("%s:%s",
		config.GetString("databases.redis.host"),
		config.GetString("databases.redis.port"),
	)
	rds.RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.GetString("databases.redis.password"),
		DB:       config.GetInt("databases.redis.db"),
	})
}
