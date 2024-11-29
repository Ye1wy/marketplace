package api

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type API struct {
	ctx context.Context
	rdb *redis.Client
}
