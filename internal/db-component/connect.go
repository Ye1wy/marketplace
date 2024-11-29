// Package db_component предоставляет функции для работы с базой данных Redis.
package db_component

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

// ConnectToRedis устанавливает соединение с сервером Redis.
// Возвращает контекст и клиент Redis.
// Если соединение не удалось установить, функция вызывает панику с ошибкой.
func ConnectToRedis() (context.Context, *redis.Client) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",	// Адрес сервера Redis.
		Password: "",	// Пароль для доступа к Redis (если требуется).
		DB:       0,	// Номер базы данных Redis.
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		slog.Error("Connection was refused")
		panic(err)
	}

	return ctx, rdb
}
