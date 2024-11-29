package data

type CacheData struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	UrlLink  string `json:"link"`
	Rating   string `json:"rating"`
	Platform string `json:"platform"`
	Category string `json:"category"`
}

type Filter struct {
	Query string `json:"query"`
	Sort  string `json:"sort"`
}
