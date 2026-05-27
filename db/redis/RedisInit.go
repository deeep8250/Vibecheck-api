package redisInit

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func RedisInit() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	ctx := context.Background()
	_, err := config.RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("redis connection failed", err.Error())
	}
	log.Println("redis connection is successful")

}
