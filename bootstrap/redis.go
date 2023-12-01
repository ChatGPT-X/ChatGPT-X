package bootstrap

import (
	"chatgpt_x/pkg/config"
	rds "chatgpt_x/pkg/redis"
	"fmt"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
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

	rds.Store = persist.NewRedisStore(rds.RDB)
	//rds.MemoryStore = persist.NewMemoryStore(86400 * time.Second)
}
