package api

import (
	"log/slog"
	db_component "marketplace/src/db-component"
	"marketplace/src/reader"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	ctx, rdb := db_component.ConnectToRedis()
	
	router := gin.Default()

	Route(router)

	router.Run("localhost:8080")
}

func Route(router *gin.Engine) {
	router.GET("/", HomePage())
	router.GET("/product", getData)
	router.POST("/add-product")
}

func getData(c *gin.Context) {
	productName := c.Query("name")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}
	
	cacheData, err := 
}

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Home page of Image Server")
	}
}

// func AddProduct() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Request.Form()
// 	}
// }
