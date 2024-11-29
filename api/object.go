// Package api предоставляет реализацию API для взаимодействия с продуктами на маркетплейсе.
package api

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// API представляет собой структуру, содержащую контекст и клиент Redis.
// Она используется для обработки запросов к API и взаимодействия с базой данных.
type API struct {
	ctx context.Context		// Контекст для управления жизненным 
	rdb *redis.Client	// Клиент Redis для доступа к базе данных.
}
