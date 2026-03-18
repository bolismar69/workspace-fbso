package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,                
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("⚠️ Redis não disponível: %v\n", err)
	} else {
		fmt.Println("⚡ Conectado ao Redis com sucesso")
	}
}

func getCacheKey(ncm, ibge string) string {
	return fmt.Sprintf("tax:calc:%s:%s", ncm, ibge)
}
