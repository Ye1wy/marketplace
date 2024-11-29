package db_component

import (
	"context"
	"marketplace/internal/data"

	"github.com/redis/go-redis/v9"
)

func Add(rdb *redis.Client, ctx context.Context, key string, data data.CacheData) {
	// TODO: Error handle

	rdb.HSet(ctx, key, data)
}
