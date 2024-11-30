package main

import "marketplace/internal/scraper"

func main() {
	scraber := scraper.NewScraber()

	scraber.GetOzon("https://www.ozon.ru/search/?text=носки&from_global=true")
	// scraber.GetOzon("https://www.ozon.ru/search/?text=bottel&from_global=true")
}
