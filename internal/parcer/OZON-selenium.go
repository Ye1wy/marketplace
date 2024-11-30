package parcer

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
	port             = 4444
)

func (scraber *Scraber) GetOzon(url string) data.CacheData {
	defer scraber.service.Stop()
	// Navigate to the Ozon website.
	// TODO: добавить динамическую ссылку
	if err := scraber.Driver.Get(url); err != nil {
		log.Fatalf("Error navigating to the website: %v", err)
	}
	time.Sleep(time.Second * 1)

	if err := scraber.Driver.Refresh(); err != nil {
		log.Fatalf("Error refresh: %s", err)
	}

	content_prefix := "//*[@id=\"paginatorContent\"]/div[1]/div/div["
	content_sufix := "]"

	for i := 0; i < 10; i++ {
		data_url := content_prefix + strconv.Itoa(i+1) + content_sufix

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
