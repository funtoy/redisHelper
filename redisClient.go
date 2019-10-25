package redisHelper

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

// "localhost:6379" "", 0
func InitClient(ip, password string, db int) {
	client = redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func Client() *redis.Client {
	return client
}
