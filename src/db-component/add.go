package db_component

import (
	"context"
	"marketplace/src/reader"

	"github.com/redis/go-redis/v9"
)

func Add(rdb *redis.Client, ctx context.Context, key string, data reader.Data) {
	rdb.HSet(ctx, key, data)
}
