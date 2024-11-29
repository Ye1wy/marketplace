package db_component

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() (context.Context, *redis.Client) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		slog.Error("Connection was refused")
		panic(err)
	}

	return ctx, rdb
}
