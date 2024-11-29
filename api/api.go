// Package api предоставляет реализацию API для взаимодействия с продуктами на маркетплейсе.
package api

import (
	"context"
	"marketplace/internal/data"
	db_component "marketplace/internal/db-component"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// RunApi запускает API сервер, устанавливает соединение с Redis и настраивает маршруты.
func RunApi() {
	ctx, rdb := db_component.ConnectToRedis()
	defer rdb.Close()

	api := NewAPI(ctx, rdb)

	router := gin.Default()
	api.Route(router)

	router.Run("localhost:8080")
}

// NewAPI создает новый экземпляр API с контекстом и клиентом Redis.
// ctx - контекст для управления жизненным циклом запросов.
// rdb - клиент Redis для доступа к базе данных.
func NewAPI(ctx context.Context, rdb *redis.Client) *API {
	return &API{ctx: ctx, rdb: rdb}
}

// Route настраивает маршруты для API.
// router - экземпляр gin.Engine, используемый для определения маршрутов.
func (api *API) Route(router *gin.Engine) {
	router.GET("/", api.HomePage())
	router.GET("/product", api.GetData)
}

// GetData обрабатывает запросы на получение данных о продукте.
// c - контекст запроса Gin.
// Возвращает JSON с данными о продукте или ошибку, если продукт не найден.
func (api *API) GetData(c *gin.Context) {
	productName := c.Query("name")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}

	products, err := db_component.ReadData(api.rdb, api.ctx, productName)
	if err != nil || products == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// TODO: parsed initialize

	data := data.CacheData{}

	db_component.Add(api.rdb, api.ctx, productName, data)

	// TODO: need to handle error from add to db function

	c.IndentedJSON(http.StatusOK, products)
}

// HomePage возвращает обработчик для главной страницы API.
// Возвращает JSON с сообщением о главной странице сервера изображений.
func (api *API) HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Home page of Image Server")
	}
}
