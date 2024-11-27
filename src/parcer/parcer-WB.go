package parcer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	UrlLink  string `json:"link"`
	Rating   string `json:"rating"`
	Platform string `json:"platform"`
	Category string `json:"category"`
}

func WriteJSON(bodyText []byte, err error) {
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}

	// Создаем файл
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	// Записываем данные в файл
	_, err = file.Write(bodyText)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	// fmt.Println("Файл успешно создан.")
}

func ParcerWB() {
	//
	// TODO: add dynamic query
	//
	//
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://search.wb.ru/exactmatch/ru/common/v7/search?ab_testing=false&appType=1&curr=rub&dest=-366541query=%D0%B4%D0%B6%D0%B8%D0%BD%D1%81%D1%8B%20%D0%B6%D0%B5%D0%BD%D1%81%D0%BA%D0%B8%D0%B5%20%D1%88%D0%B8%D1%80%D0%BE%D0%BA%D0%B8%D0%B5&resultset=filters&spp=30&suppressSpellcheck=false", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Set("origin", "https://www.wildberries.ru")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.wildberries.ru/catalog/0/search.aspx?search=%D0%B4%D0%B6%D0%B8%D0%BD%D1%81%D1%8B%20%D0%B6%D0%B5%D0%BD%D1%81%D0%BA%D0%B8%D0%B5%20%D1%88%D0%B8%D1%80%D0%BE%D0%BA%D0%B8%D0%B5")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Set("x-captcha-id", "Catalog 1|1|1733812057|AA==|783f1311c88144ff9f06ff6a1f7d70c6|mLrGEJvghWJUJXC4tMl3BiCX6BPRRiI2wICV56GXcO2")
	req.Header.Set("x-queryid", "qid783729000173260245520241126062850")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	WriteJSON(bodyText, err)

}
