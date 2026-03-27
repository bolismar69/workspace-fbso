package cache // <--- Adicione isto

import (
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
