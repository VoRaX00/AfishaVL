package domain

import "time"

type RefreshToken struct {
	Id               int       `json:"id" db:"id"`
	UserId           int       `json:"user_id" db:"user_id"`
	RefreshTokenHash string    `json:"refresh_token_hash" db:"refresh_token_hash"`
	Ip               string    `json:"ip" db:"ip"`
	ExpiresAt        time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}
