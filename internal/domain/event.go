package domain

type Event struct {
	Id           int    `json:"id" db:"id"`
	UserId       int    `json:"userId" db:"user_id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	CityId       int    `json:"cityId" db:"city_id"`
	CategoryId   int    `json:"categoryId" db:"category_id"`
	Image        byte   `json:"image" db:"image"`
	PriceId      int    `json:"priceId" db:"price_id"`
	AgeLimit     int    `json:"ageLimit" db:"age_limit"`
	ScheduleId   int    `json:"scheduleId" db:"schedule_id"`
	Registration bool   `json:"registration" db:"registration"`
	ContactId    int    `json:"contactId" db:"contact_id"`
}
