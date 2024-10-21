package infrastructure

import (
	"Afisha/internal/domain"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Create(token domain.RefreshToken) error {
	query := "INSERT INTO refresh_tokens (user_id, refresh_token_hash, ip, expires_at) VALUES ($1, $2, $3, $4)"
	err := r.db.QueryRow(query, token.UserId, token.RefreshTokenHash, token.Ip, token.ExpiresAt)
	return err.Err()
}

func (r *TokenRepository) Update(newTokenHash, tokenHash, ipClient string, ttl time.Duration) (string, error) {
	var userId string
	var expiresAt time.Time
	var ip string
	query := "SELECT user_id, expires_at, ip FROM refresh_tokens WHERE refresh_token_hash = $1"

	err := r.db.QueryRow(query, tokenHash).Scan(&userId, &expiresAt, &ip)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("refresh token not found")
		}
		return "", err
	}

	if time.Now().After(expiresAt) {
		return "", errors.New("refresh token expired")
	}

	expiresAt = time.Now().Add(ttl)
	query = "UPDATE refresh_tokens SET refresh_token_hash = $1, expires_at = $2, ip = $3, created_at=CURRENT_TIMESTAMP WHERE user_id = $4"
	_, err = r.db.Exec(query, newTokenHash, expiresAt, ipClient, userId)
	if ip != ipClient {
		return ipClient, nil
	}
	return "", err
}
