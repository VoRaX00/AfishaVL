package domain

import "time"

type UserToLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRegister struct {
	Phone    string
	Password string
}

type CreateEvent struct {
	UserId           string
	NameEvent        string
	CityId           int
	CategoryId       int
	PosterImg        byte
	AgeLimit         int
	DiscriptionEvent string
	Price            int
	Registration     bool
	Comment          string
	Place            string
	Address          string
	DateStart        time.Time
	TimeStart        time.Time
	ContactId        int
}
