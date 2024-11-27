package db_component

import (
	"context"
	"log/slog"
	"marketplace/src/reader"

	"github.com/redis/go-redis/v9"
)

func ReadData(rdb *redis.Client, ctx context.Context, key string) (*reader.Data, error) {
	var data reader.Data

	err := rdb.HGetAll(ctx, key).Scan(&data)

	if err != nil {
		slog.Error("Cannot find data in database")
		return nil, err
	}

	return &data, nil
}
