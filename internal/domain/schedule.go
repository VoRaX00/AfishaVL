package domain

import "time"

type Schedule struct {
	Id      int       `json:"id" db:"id"`
	Place   string    `json:"place" db:"place"`
	Address string    `json:"address" db:"address"`
	Dates   time.Time `json:"dates" db:"dates"`
}
