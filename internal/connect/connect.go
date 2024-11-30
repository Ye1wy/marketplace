package connect

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() {
	ctx := context.Background()

	rbd := redis.NewClient(&redis.Options{
		Addr: localhost,
	})
}
