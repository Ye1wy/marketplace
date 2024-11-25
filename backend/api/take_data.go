package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Создаем новый парсер
	c := colly.NewCollector(
		colly.AllowedDomains("ozon.ru"),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		r.Headers.Set("Accept-Language", "ru-RU,ru;q=0.9")
	})

	// Обрабатываем карточку продукта
	// c.OnHTML("div.j65_23.j75_23.tile-root", func(e *colly.HTMLElement) {
	// 	// Извлекаем цену
	// 	price := e.ChildText("span.c3022-a1.tsHeadline500Medium.c3022-b8.c3022-a6")
	// 	price = strings.ReplaceAll(price, "\u00a0", "") // Убираем неразрывные пробелы
	// 	fmt.Println("Цена:", price)
	// })

	// Достаем всю страницу
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.Text) // Вывод всего содержимого
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Ошибка запроса:", err)
		fmt.Println("Статус:", r.StatusCode)
	})

	// Переход на страницу с продуктами
	err := c.Visit("https://ozon.ru/search?from_global=true&sorting=rating&text=бутылка")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/chromedp/chromedp"
// )

// func main() {
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Таймаут
// 	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
// 	defer cancel()

// 	var html string
// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate("https://ozon.ru/search?from_global=true&sorting=rating&text=бутылка"),
// 		chromedp.WaitVisible(`div.j65_23.j75_23.tile-root`), // Ждем загрузки контента
// 		chromedp.OuterHTML(`html`, &html),
// 	)
// 	if err != nil {
// 		fmt.Println("Ошибка:", err)
// 		return
// 	}

// 	fmt.Println("HTML страницы:", html)
// }
