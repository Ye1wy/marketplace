package db_component

import (
	"context"
	"log/slog"
	"marketplace/src/data"

	"github.com/redis/go-redis/v9"
)

func ReadData(rdb *redis.Client, ctx context.Context, key string) (*data.CacheData, error) {
	var data data.CacheData

	err := rdb.HGetAll(ctx, key).Scan(&data)

	if err != nil {
		slog.Error("Cannot find data in database")
		return nil, err
	}

	return &data, nil
}
