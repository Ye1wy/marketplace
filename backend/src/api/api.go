package main

import (
	"log/slog"
	"marketplace/backend/src/reader"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	Route(router)

	router.Run("localhost:8080")
}

func Route(router *gin.Engine) {
	router.GET("/", HomePage())
	router.GET("/product", getData)
}

func getData(c *gin.Context) {
	file, err := os.Open("storage/Data.json")

	if err != nil {
		slog.Info("File not found")
		c.Status(http.StatusNotFound)
		return
	}

	defer file.Close()

	var data reader.Data

	err = data.Parse(file)

	c.IndentedJSON(http.StatusOK, data)
}

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Home page of Image Server")
	}
}
