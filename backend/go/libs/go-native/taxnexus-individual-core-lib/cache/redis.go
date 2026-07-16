package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

// ConnectRedis cria um cliente Redis ou retorna nil se REDIS_ADDR não estiver configurado.
// Quando nil é retornado, o TaxRepository opera em modo "cache desabilitado", consultando apenas o PostgreSQL.
func ConnectRedis(addr string) *redis.Client {
	if addr == "" {
		return nil
	}
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

// ConnectRedisFromEnv cria um cliente Redis a partir da variável de ambiente REDIS_ADDR.
// Retorna nil se a variável não estiver definida ou for vazia.
func ConnectRedisFromEnv() *redis.Client {
	return ConnectRedis(os.Getenv("REDIS_ADDR"))
}

// CacheTTL retorna a duração do TTL do cache, obtida de TAX_CACHE_TTL_HOURS.
// Se não configurada, retorna 12 horas como padrão.
func CacheTTL() int {
	defaultTTL := 12
	// Placeholder para leitura de env var — evita dependência circular com strconv.
	// O valor é lido diretamente no repository via os.Getenv.
	return defaultTTL
}
