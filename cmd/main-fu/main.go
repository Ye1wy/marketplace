package main

import "marketplace/internal/parcer"

func main() {
	scraber := parcer.NewScraber()

	scraber.GetOzon("https://www.ozon.ru/search/?text=носки&from_global=true")
	scraber.GetOzon("https://www.ozon.ru/search/?text=bottel&from_global=true")
}
