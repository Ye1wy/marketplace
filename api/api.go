package api

import (
	"context"
	"log/slog"
	db_component "marketplace/internal/db-component"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type API struct {
	ctx context.Context
	rdb *redis.Client
}

func NewAPI(ctx context.Context, rdb *redis.Client) *API {
	return &API{ctx: ctx, rdb: rdb}
}

func (api *API) Run() {
	router := gin.Default()
	api.Route(router)

	router.Run("localhost:8080")
}

func (api *API) Route(router *gin.Engine) {
	router.GET("/", homePage())
	router.POST("/product", productHandler(api))
}

func homePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Home page of Image Server")
	}
}

func productHandler(api *API) gin.HandlerFunc {
	return func(c *gin.Context) {
		productName := c.Query("name")
		if productName == "" {
			slog.Info("No product name in request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
			return
		}

		products, err := db_component.ReadData(api.rdb, api.ctx, productName)
		if err != nil || products == nil {
			if err := api.rdb.Publish(api.ctx, "api_to_scraper", productName).Err(); err != nil {
				slog.Info("Failde to publish message")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message"})
				return
			}

			reply := api.rdb.Subscribe(api.ctx, "sort_to_api")
			defer reply.Close()

			msg := <-reply.Channel()

			// cacheData, err := waitForMessage(api.ctx, reply)
			cacheData, err := db_component.ConvertToJASON(msg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
				return
			}

			db_component.Add(api.rdb, api.ctx, productName, cacheData)
			c.JSON(http.StatusOK, cacheData)
			return
		}

		c.IndentedJSON(http.StatusOK, products)
	}
}

// func waitForMessage(ctx context.Context, pubsub *redis.PubSub) (*data.CacheData, error) {
// 	// Создаем канал для сообщений
// 	messageChan := pubsub.Channel()

// 	for {
// 		select {
// 		case msg := <-messageChan: // Ждем новое сообщение
// 			if msg == nil {
// 				continue // Если канал пустой, пропускаем
// 			}

// 			// // Конвертируем сообщение в структуру CacheData
// 			// cacheData, err := db_component.ConvertToJASON(msg)
// 			// if err != nil {
// 			// 	return nil, err
// 			// }
// 			// return &cacheData, nil

// 			// case <-ctx.Done(): // Завершаем работу, если контекст отменен
// 			// 	return nil, ctx.Err()
// 		}
// 	}
// }

// func sortRatingHandler(api *API) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		productName := ctx.Query("name")
// 		if productName == "" {
// 			slog.Info("No product name in request")
// 			ctx.JSON(http.StatusOK, gin.H{"error": "Product name is required"})
// 			return
// 		}

// 		products, _ := db_component.ReadData(api.rdb, api.ctx, productName)

// 	}
// }
