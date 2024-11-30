package scraper

import (
	"fmt"
	"log"
	"log/slog"
	"marketplace/internal/data"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Scraber struct {
	service *selenium.Service
	Driver  selenium.WebDriver
}

func NewScraber() *Scraber {
	new_service, err := selenium.NewChromeDriverService(chromeDriverPath, port)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
		return nil
	}

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{"--headles-new"}})

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

	return &Scraber{service: new_service, Driver: driver}
}

type ScrabData interface {
	GetOzon(url string) data.CacheData
	GetWb(url string) data.CacheData
}
