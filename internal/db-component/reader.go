// Package db_component предоставляет функции для работы с базой данных Redis.
package db_component

import (
	"context"
	"log/slog"
	"marketplace/internal/data"

	"github.com/redis/go-redis/v9"
)

// ReadData извлекает данные из Redis по указанному ключу.
// rdb - клиент Redis, используемый для выполнения операций с базой данных.
// ctx - контекст для управления жизненным циклом операции.
// key - уникальный ключ, по которому будут извлекаться данные.
// Возвращает указатель на структуру CacheData и ошибку, если произошла ошибка при извлечении данных.
func ReadData(rdb *redis.Client, ctx context.Context, key string) (*data.CacheData, error) {
	var data data.CacheData

	err := rdb.HGetAll(ctx, key).Scan(&data)

	if err != nil {
		slog.Error("Cannot find data in database")
		return nil, err
	}

	return &data, nil
}
