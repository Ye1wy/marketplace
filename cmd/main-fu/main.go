package main

import (
	"marketplace/internal/data"
	"marketplace/internal/scraper"
)

func main() {
	scrap := scraper.NewScraper()
	cacheWB := data.CacheData{
		Products: make([]data.Product, 10),
		Request:  "",
		Sort:     "default",
	}
	wildberriesScraper := scraper.NewWildberries(scrap)
	url := "https://www.wildberries.ru/catalog/0/search.aspx?search=bottle"
	scraper.Navigate(wildberriesScraper, url)
	scraper.ScrabElements(wildberriesScraper, &cacheWB)
	scraper.ScrabUrl(wildberriesScraper, &cacheWB)
	scraper.ScrabImg(wildberriesScraper, &cacheWB)
}
