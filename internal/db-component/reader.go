package db_component

import (
	"marketplace/internal/data"
	"strings"

	"github.com/redis/go-redis/v9"
)

func RedisToCacheData(cmd *redis.MapStringStringCmd) (data.CacheData, error) {
	res, err := cmd.Result()
	if err != nil {
		return data.CacheData{}, err
	}

	cacheData := data.CacheData{
		Products: []data.Product{},
	}

	productMap := make(map[string]*data.Product)

	// Разбираем мапу
	for key, value := range res {
		parts := strings.Split(key, ":")
		if len(parts) < 3 {
			continue // Пропускаем некорректные ключи
		}

		// product:0:name -> parts[0]="product", parts[1]="0", parts[2]="name"
		productIndex := parts[1]
		field := parts[2]

		// Получаем или создаем продукт
		if _, exists := productMap[productIndex]; !exists {
			productMap[productIndex] = &data.Product{}
		}
		product := productMap[productIndex]

		// Заполняем поле
		switch field {
		case "name":
			product.Name = value
		case "price":
			product.Price = value
		case "link":
			product.UrlLink = value
		case "rating":
			product.Rating = value
		case "image":
			product.Image = value
		}
	}

	// Переносим продукты из map в слайс
	for _, product := range productMap {
		cacheData.Products = append(cacheData.Products, *product)
	}

	// Добавляем поля Request и Sort, если есть
	if req, ok := res["request"]; ok {
		cacheData.Request = req
	}
	if sort, ok := res["sort"]; ok {
		cacheData.Sort = sort
	}

	return cacheData, nil
}
