// Package data предоставляет структуры для работы с данными продуктов на маркетплейсе.
package data

// CacheData представляет собой структуру для хранения данных о продуктах в кэше.
type CacheData struct {
	Products []Product `json:"products"` // Список продуктов.
}

// Product представляет собой структуру, описывающую продукт на маркетплейсе.
type Product struct {
	Id       string `json:"id"` // Уникальный идентификатор продукта.
	Name     string `json:"name"` // Название продукта.
	Price    string `json:"price"` // Цена продукта.
	UrlLink  string `json:"link"` // URL-ссылка на продукт.
	Rating   string `json:"rating"` // Рейтинг продукта.
	Platform string `json:"platform"` // Платформа, на которой доступен продукт.
	Category string `json:"category"` // Категория продукта.
}

// Filter представляет собой структуру для фильтрации продуктов.
type Filter struct {
	Query string `json:"query"` // Запрос для фильтрации продуктов по имени или другим характеристикам.
	Sort  string `json:"sort"` // Параметр сортировки (например, по цене или рейтингу).
}
