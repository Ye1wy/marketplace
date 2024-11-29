// Package db_component предоставляет функции для взаимодействия с базой данных, в частности, для работы с Redis.
package db_component

import (
	"context"
	"marketplace/internal/data"

	"github.com/redis/go-redis/v9"
)

// Add добавляет данные о продукте в Redis по указанному ключу.
// rdb - клиент Redis, используемый для выполнения операций с базой данных.
// ctx - контекст для управления жизненным циклом операции.
// key - уникальный ключ, под которым будут храниться данные в Redis.
// data - структура CacheData, содержащая информацию о продуктах, которую необходимо сохранить.
func Add(rdb *redis.Client, ctx context.Context, key string, data data.CacheData) {
	// TODO: Error handle

	rdb.HSet(ctx, key, data)
}
