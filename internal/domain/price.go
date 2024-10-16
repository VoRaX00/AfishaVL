package domain

type Price struct {
	Id        int    `json:"id" db:"id"`
	PriceFrom int    `json:"priceFrom" db:"price_from"`
	PriceTo   int    `json:"priceTo" db:"price_to"`
	Comment   string `json:"comment" db:"comment"`
	Free      bool   `json:"free" db:"free"`
}
