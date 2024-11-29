package parcer

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	seleniumPath     = "/usr/local/bin/selenium/selenium-server-standalone-3.9.0.jar"
	chromeDriverPath = "./chromedriver"
	port             = 4444
)

func ScrapingOZON() {
	opts := []selenium.ServiceOption{
		selenium.Output(nil), // Discard output.
	}
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{"--headles-new"}})

	// Connect to the WebDriver instance running locally.
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	if err != nil {
		log.Fatalf("Error opening session: %s", err)
	}
	defer wd.Quit()

	// Navigate to the Ozon website.
	// TODO: добавить динамическую ссылку
	if err := wd.Get("https://www.ozon.ru/category/odezhda-obuv-i-aksessuary-7500/?category_was_predicted=true&deny_category_prediction=true&from_global=true&text=%D0%BD%D0%BE%D1%81%D0%BA%D0%B8"); err != nil {
		log.Fatalf("Error navigating to the website: %v", err)
	}
	time.Sleep(time.Second * 1)

	if err := wd.Refresh(); err != nil {
		log.Fatalf("Error refresh: %s", err)
	}

	time.Sleep(time.Second * 5)

	// text, _ := wd.PageSource()
	// fmt.Printf("%s", text)

	// Use XPath to find elements and gather data.
	// elem, err := wd.FindElement(selenium.ByXPATH, "//h1[@data-index='0']")
	elem, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"paginatorContent\"]/div[1]/div/div[1]")
	if err != nil {
		log.Fatalf("Error finding element: %v", err)
	}

	// Extract text from the element.

	text, err := elem.Text()

	if err != nil {
		log.Fatalf("Error getting text: %v", err)
	}

	fmt.Printf("Extracted text: %s\n", text)
}
