package domain

import "time"

type RefreshToken struct {
	Id               int       `json:"id" db:"id"`
	UserId           int       `json:"userId" db:"user_id"`
	RefreshTokenHash string    `json:"refreshTokenHash" db:"refresh_token_hash"`
	Ip               string    `json:"ip" db:"ip"`
	ExpiresAt        time.Time `json:"expiresAt" db:"expires_at"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}
