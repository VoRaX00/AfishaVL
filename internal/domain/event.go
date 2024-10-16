package domain

type Event struct {
	Id           int    `json:"id" db:"id"`
	UserId       int    `json:"user_id" db:"user_id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	CityId       int    `json:"city_id" db:"city_id"`
	CategoryId   int    `json:"category_id" db:"category_id"`
	Image        byte   `json:"image" db:"image"`
	PriceId      int    `json:"price_id" db:"price_id"`
	AgeLimit     int    `json:"age_limit" db:"age_limit"`
	ScheduleId   int    `json:"schedule_id" db:"schedule_id"`
	Registration bool   `json:"registration" db:"registration"`
	ContactId    int    `json:"contact_id" db:"contact_id"`
}
