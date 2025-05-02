package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tranduckhuy/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {

	config := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
		PoolSize: config.PoolSize, // set connection pool size
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Fatal("Redis connection failed: ", zap.Error(err))
	}

	global.Logger.Info("Redis connection successful")
	global.Redis = rdb

	redisExample()
}

func redisExample() {
	err := global.Redis.Set(ctx, "score", 100, 10*time.Second).Err()

	if err != nil {
		fmt.Println("Error setting value in Redis:", err)
		return
	}

	val, err := global.Redis.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error getting value from Redis:", err)
		return
	}

	global.Logger.Info("Redis value", zap.String("score", val))
}
