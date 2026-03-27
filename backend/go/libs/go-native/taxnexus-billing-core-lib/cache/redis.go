// path: backend/go/libs/go-native/taxnexus-billing-core-lib/cache/redis.go
package cache

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func ConnectRedis(addr string) *redis.Client {
	slog.Info("Inicializando cache Redis para alíquotas", "host", addr)
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}


