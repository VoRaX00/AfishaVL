package domain

type Price struct {
	Id        int    `json:"id" db:"id"`
	PriceFrom int    `json:"price_from" db:"price_from"`
	PriceTo   int    `json:"price_to" db:"price_to"`
	Comment   string `json:"comment" db:"comment"`
	Free      bool   `json:"free" db:"free"`
}
