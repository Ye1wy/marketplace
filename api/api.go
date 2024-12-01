package api

import (
	"context"
	"log/slog"
	"marketplace/internal/data"
	db_component "marketplace/internal/db-component"
	"marketplace/internal/scraper"
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
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
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
		slog.Info("Received POST /product request",
			"productName", productName,
			"queryParams", c.Request.URL.Query())

		if productName == "" {
			slog.Warn("No product name in request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
			return
		}

		// products, err := db_component.ReadData(api.rdb, api.ctx, productName)

		// if err != nil {
		slog.Info("Product not found in cache; publishing to scraper",
			"productName", productName)

		cacheWB := data.CacheData{
			Products: make([]data.Product, 10),
			Request:  productName,
			Sort:     "default",
		}

		scrap := scraper.NewScraper()
		wildberriesScraper := scraper.NewWildberries(scrap)
		url := "https://www.wildberries.ru/catalog/0/search.aspx?search=" + productName
		scraper.Navigate(wildberriesScraper, url)
		scraper.ScrabElements(wildberriesScraper, &cacheWB)
		scraper.ScrabUrl(wildberriesScraper, &cacheWB)
		scraper.ScrabImg(wildberriesScraper, &cacheWB)

		cacheOzon := data.CacheData{
			Products: make([]data.Product, 10),
			Request:  productName,
			Sort:     "default",
		}

		scrap.Quit()

		scrap = scraper.NewScraper()

		ozonScraper := scraper.NewOzon(scrap)
		url = "https://www.ozon.ru/search/?text=" + productName + "&from_global=true"
		scraper.Navigate(ozonScraper, url)
		scraper.ScrabElements(ozonScraper, &cacheOzon)
		scraper.ScrabUrl(ozonScraper, &cacheOzon)
		scraper.ScrabImg(ozonScraper, &cacheOzon)
		scrap.Quit()

		products := data.MergeCacheData(cacheWB, cacheOzon)
		db_component.Add(api.rdb, api.ctx, productName, products)
		c.JSON(http.StatusOK, products)
	}

	// slog.Info("Returning cached product data",
	// 	"productName", productName,
	// 	"data", products)
	// c.IndentedJSON(http.StatusOK, products)
	// // }
}
