package data

type CacheData struct {
	Products []Product `json:"products"`
	Request  string    `json:"request"`
	Sort     string    `json:"sort"`
}

type Product struct {
	Name    string `json:"name"`
	Price   string `json:"price"`
	UrlLink string `json:"link"`
	Rating  string `json:"rating"`
	Image   string `json:"image"`
}

func MergeCacheData(data1, data2 CacheData) CacheData {
	mergedData := CacheData{
		Products: make([]Product, 0, len(data1.Products)+len(data2.Products)),
		Request:  data1.Request, // Сохраняем запрос из первого объекта
		Sort:     data1.Sort,    // Сохраняем сортировку из первого объекта
	}

	// Добавляем все продукты из первого объекта
	productMap := make(map[string]struct{}) // Для уникальности продуктов по URL
	for _, product := range data1.Products {
		if product.UrlLink != "" {
			productMap[product.UrlLink] = struct{}{}
		}
		mergedData.Products = append(mergedData.Products, product)
	}

	// Добавляем продукты из второго объекта, проверяя на дубликаты
	for _, product := range data2.Products {
		if _, exists := productMap[product.UrlLink]; !exists {
			if product.UrlLink != "" {
				productMap[product.UrlLink] = struct{}{}
			}
			mergedData.Products = append(mergedData.Products, product)
		}
	}

	return mergedData
}
