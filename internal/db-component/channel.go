package db_component

import (
	"encoding/json"
	"log/slog"
	"marketplace/internal/data"

	"github.com/redis/go-redis/v9"
)

func ConvertToJASON(msg *redis.Message) (data.CacheData, error) {
	var cacheData data.CacheData

	// Декодируем строку сообщения из канала в структуру CacheData
	if err := json.Unmarshal([]byte(msg.Payload), &cacheData); err != nil {
		slog.Error("Failed to unmarshal Redis message into CacheData")
		return data.CacheData{}, err
	}

	return cacheData, nil
}
