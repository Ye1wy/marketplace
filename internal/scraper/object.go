package scraper

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverTruePath = "./internal/chromedriver/chromedriver"
	// chromeDriverPath = "./internal/chromedriver/chromedriver-mac"
	port = 8080
)

type Scraper struct {
	Service *selenium.Service
	Driver  selenium.WebDriver
}

func NewScraper() *Scraper {
	new_service, err := selenium.NewChromeDriverService(chromeDriverTruePath, port)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
		return nil
	}

	// caps := selenium.Capabilities{}
	caps := selenium.Capabilities{"browserName": "chrome", "goog:chromeOptions": map[string]interface{}{"args": []string{"--headless", "--disable-gpu", "--no-sandbox", "--disable-dev-shm-usage"}}}
	// caps.AddChrome(chrome.Capabilities{Args: []string{"--headless"}})

	// Connect to the WebDriver instance running locally.
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error opening session: %s", err)
		return nil
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		slog.Error("Error: cannot maximize window")
		return nil
	}

	return &Scraper{Service: new_service, Driver: driver}
}

func (scraper *Scraper) Quit() {
	scraper.Driver.Quit()
	scraper.Service.Stop()
}

type ScrapingConfig struct {
	ContentPrefix string
	ContentSuffix string
}

type Ozon struct {
	Scraper *Scraper
	Config  map[string]ScrapingConfig
}

func NewOzon(scraper *Scraper) *Ozon {
	return &Ozon{
		Scraper: scraper,
		Config: map[string]ScrapingConfig{
			"elements": {ContentPrefix: "//*[@id=\"paginatorContent\"]/div[1]/div/div[", ContentSuffix: "]"},
			"url":      {ContentPrefix: "//*[@id=\"paginatorContent\"]/div[1]/div/div[1]/div/a", ContentSuffix: ""},
			"images":   {ContentPrefix: "//*[@id=\"paginatorContent\"]/div[1]/div/div[1]/div/a/div/div[1]/img", ContentSuffix: ""},
		},
	}
}

type Wildberries struct {
	Scraper *Scraper
	Config  map[string]ScrapingConfig
}

func NewWildberries(scraper *Scraper) *Wildberries {
	return &Wildberries{
		Scraper: scraper,
		Config: map[string]ScrapingConfig{
			"elements": {ContentPrefix: "/html/body/div[1]/main/div[2]/div[2]/div[2]/div/div/div[4]/div[1]/div[1]/div/article[", ContentSuffix: "]"},
			"url":      {ContentPrefix: "/html/body/div[1]/main/div[2]/div[2]/div[2]/div/div/div[4]/div[1]/div[1]/div/article[", ContentSuffix: "]/div/a"},
			"images":   {ContentPrefix: "/html/body/div[1]/main/div[2]/div[2]/div[2]/div/div/div[4]/div[1]/div[1]/div/article[", ContentSuffix: "]/div/div[1]/div[2]/img[1]"},
		},
	}
}

type ScraperInterface interface {
	GetConfig(key string) ScrapingConfig
	GetDriver() selenium.WebDriver
}
