package api

import (
	"context"
	"log/slog"
	"marketplace/internal/data"
	db_component "marketplace/internal/db-component"
	"marketplace/internal/scraper"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type API struct {
	ctx context.Context
	rdb *redis.Client
}

type PageData struct {
	Message string
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
	router.GET("/", homePage(router))
	router.GET("/product", productHandler(api))
}

func homePage(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("./index.html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Hello, World!",
		})
		router.Static("./css/styles.css", "./index.html")
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

		products := data.MergeCacheData(cacheWB, cacheOzon)
		db_component.Add(api.rdb, api.ctx, productName, products)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("static/index.html"))
			data := PageData{
				Message: "Hello, this is dynamic data!",
			}
			tmpl.Execute(w, data)
		})
		c.JSON(http.StatusOK, products)
	}

	// slog.Info("Returning cached product data",
	// 	"productName", productName,
	// 	"data", products)
	// c.IndentedJSON(http.StatusOK, products)
	// // }
}
