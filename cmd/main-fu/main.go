package main

import "marketplace/internal/scraper"

func main() {
	scrap := scraper.NewScraper()

	wildberriesScraper := scraper.NewWildberries(scrap)
	url := "https://www.wildberries.ru/catalog/0/search.aspx?search=bottle"
	scraper.Navigate(wildberriesScraper, url)
	scraper.ScrabElements(wildberriesScraper)
	scraper.ScrabUrl(wildberriesScraper)
	scraper.ScrabImg(wildberriesScraper)
}
