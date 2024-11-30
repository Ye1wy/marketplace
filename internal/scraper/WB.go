package scraper

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

const (
	limit = 10
)

func (dataScraper *Wildberries) Navigate(url string) {
	if err := dataScraper.Scraper.Driver.Get(url); err != nil {
		slog.Error("Unviable url in WB driver navigate component")
	}
}

func (dataScraper *Wildberries) ScrabElements() []string {
	key := "elements"
	collectedData := make([]string, limit)

	time.Sleep(time.Second * 6)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.Config[key].ContentPrefix + strconv.Itoa(i+1) + dataScraper.Config[key].ContentSuffix
		element, err := dataScraper.Scraper.Driver.FindElement(selenium.ByXPATH, htmlCase)
		fmt.Println(htmlCase)

		if err != nil {
			slog.Error("Not find element in WB driver")
			continue
		}

		text, err := element.Text()
		if err != nil {
			slog.Error("Error in text extracting")
			continue
		}

		fmt.Println(text)
	}

	return collectedData
}

func (dataScraper *Wildberries) ScrabUrl() []string {
	key := "url"
	collectedData := make([]string, limit)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.Config[key].ContentPrefix + strconv.Itoa(i+1) + dataScraper.Config[key].ContentSuffix
		url, err := dataScraper.Scraper.Driver.FindElement(selenium.ByXPATH, htmlCase)
		if err != nil {
			slog.Error("Error in get product card url from html case in WB driver")
			continue
		}

		fmt.Println(url)
	}

	return collectedData
}

func (dataScraper *Wildberries) ScrabImg() []string {
	key := "image"
	collectedData := make([]string, limit)

	for i := 0; i < limit; i++ {
		htmlCase := dataScraper.Config[key].ContentPrefix + strconv.Itoa(i+1) + dataScraper.Config[key].ContentSuffix
		image, err := dataScraper.Scraper.Driver.FindElement(selenium.ByXPATH, htmlCase)
		if err != nil {
			slog.Error("Error in get image from html case in WB driver")
			continue
		}

		fmt.Println(image)
	}

	return collectedData
}
