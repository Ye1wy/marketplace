package api

import (
	"context"
	"marketplace/internal/data"
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
}
