package scraper

import (
	"fmt"
	"log"
	"log/slog"
	"marketplace/internal/data"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "./internal/chromedriver/chromedriver"
	// chromeDriverPath = "./internal/chromedriver/chromedriver-mac"
	port = 4444
)

func ScrabImg(scraber *Scraber) data.CacheData {
	content_prefix := "//*[@id=\"paginatorContent\"]/div[1]/div/div[1]/div/a/div/div[1]/img"
	// content_sufix := "]"
	elem, err := scraber.Driver.FindElement(selenium.ByXPATH, content_prefix)

	if err != nil {
		log.Fatalf("Error getting current URL: %v", err)
	}

	img, err := elem.GetAttribute("src")
	if err != nil {
		slog.Info("Err element in get atribut")
	}
	// TODO: did whrite json
	fmt.Printf(" %s \n", img)
	return data.CacheData{}
}

func ScrabElements(scraber *Scraber) data.CacheData {
	content_prefix := "//*[@id=\"paginatorContent\"]/div[1]/div/div["
	content_sufix := "]"

	for i := 0; i < 10; i++ {
		data_url := content_prefix + strconv.Itoa(i+1) + content_sufix
		ScrabUrl(scraber)
		ScrabImg(scraber)

		elem, err := scraber.Driver.FindElement(selenium.ByXPATH, data_url)
		if err != nil {
			slog.Error("Not find element")
		}

		text, err := elem.Text()
		if err != nil {
			slog.Error("Error in text extracting")
			continue
		}

		// Extract text from the element.
		fmt.Println(text)
	}

	return data.CacheData{}
}

func ScrabUrl(scraber *Scraber) data.CacheData {
	content_prefix := "//*[@id=\"paginatorContent\"]/div[1]/div/div[1]/div/a"
	//*[@id="paginatorContent"]/div[1]/div/div[1]/div/a

	elem, err := scraber.Driver.FindElement(selenium.ByXPATH, content_prefix)

	url, err := elem.GetAttribute("href")
	if err != nil {
		slog.Info("Err element in get atribut")
	}

	fmt.Printf(" %s \n", url)

	return data.CacheData{}
}

func (scraber *Scraber) GetOzon(url string) {
	// defer scraber.service.Stop()

	if err := scraber.Driver.Get(url); err != nil {
		log.Fatalf("Error navigating to the website: %v", err)
	}
	time.Sleep(time.Second * 1)

	if err := scraber.Driver.Refresh(); err != nil {
		log.Fatalf("Error refresh: %s", err)
	}

	ScrabElements(scraber)
	ScrabUrl(scraber)
	ScrabImg(scraber)

}
