package scraper

import (
	"fmt"
	"log/slog"
	"marketplace/internal/data"
	"regexp"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

const (
	limit = 10
)

func Navigate(dataScraper ScraperInterface, url string) {
	if err := dataScraper.GetDriver().Get(url); err != nil {
		slog.Error("Unviable url in WB driver navigate component")
	}

	time.Sleep(time.Second * 5)
}

func ScrabElements(dataScraper ScraperInterface) []string {
	key := "elements"
	collectedData := make([]string, limit)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.GetConfig(key).ContentPrefix + strconv.Itoa(i+1) + dataScraper.GetConfig(key).ContentSuffix
		element, err := dataScraper.GetDriver().FindElement(selenium.ByXPATH, htmlCase)

		if err != nil {
			slog.Error("ScrabElements: not find element in WB driver")
			continue
		}

		text, err := element.Text()
		if err != nil {
			slog.Error("ScrabElement: error in text extracting")
			continue
		}

		fmt.Println(text)
	}

	return collectedData
}

func ScrabUrl(dataScraper ScraperInterface) []string {
	key := "url"
	collectedData := make([]string, limit)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.GetConfig(key).ContentPrefix + strconv.Itoa(i+1) + dataScraper.GetConfig(key).ContentSuffix
		element, err := dataScraper.GetDriver().FindElement(selenium.ByXPATH, htmlCase)
		if err != nil {
			slog.Error("ScragUrl: error in find element in html case in WB driver")
			continue
		}

		url, err := element.GetAttribute("href")
		if err != nil {
			slog.Error("ScrabUrl: error in get url from html case in WB driver")
		}

		fmt.Println(url)
	}

	return collectedData
}

func ScrabImg(dataScraper ScraperInterface) []string {
	key := "images"
	collectedData := make([]string, limit)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.GetConfig(key).ContentPrefix + strconv.Itoa(i+1) + dataScraper.GetConfig(key).ContentSuffix
		element, err := dataScraper.GetDriver().FindElement(selenium.ByXPATH, htmlCase)
		if err != nil {
			slog.Error("ScrabImg: error in finding element")
			continue
		}

		image, err := element.GetAttribute("src")
		if err != nil {
			slog.Error("ScrabImg: error in get image from html case")
			continue
		}

		fmt.Println(image)
	}

	return collectedData
}

func WriteElementToJson(text string, data data.Product) data.Product {
	priceRegex := regexp.MustCompile(`\d{1,3}(?:\s\d{3})? ₽`)   // Цены
	ratingRegex := regexp.MustCompile(`\d+(?:,\d+)?`)           // Рейтинги
	titleRegex := regexp.MustCompile(`[А-Яа-яA-Za-z0-9 /]+ мл`) // Названия с "мл"
	// deliveryRegex := regexp.MustCompile(`(?:Завтра|[0-9]+ [а-я]+)`) // Доставка (Завтра или дата)

	// Найти все совпадения
	prices := priceRegex.FindAllString(text, -1)
	ratings := ratingRegex.FindAllString(text, -1)
	titles := titleRegex.FindAllString(text, -1)
	// deliveries := deliveryRegex.FindAllString(text, -1)

	if len(prices) > 0 {
		data.Price = prices[0]
	}

	if len(ratings) > 0 {
		data.Rating = ratings[0]
	}

	if len(titles) > 0 {
		data.Name = titles[0]
	}

	return data
}

func (ozon *Ozon) GetConfig(key string) ScrapingConfig {
	return ozon.Config[key]
}

func (ozon *Ozon) GetDriver() selenium.WebDriver {
	return ozon.Scraper.Driver
}

func (wildberries *Wildberries) GetConfig(key string) ScrapingConfig {
	return wildberries.Config[key]
}

func (wildberries *Wildberries) GetDriver() selenium.WebDriver {
	return wildberries.Scraper.Driver
}
