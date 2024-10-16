package infrastructure

import "github.com/jmoiron/sqlx"

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Create(userId, ip, token string) error {
	return nil
}

func (r *TokenRepository) Update(userId, ip, token string) error {
	return nil
}
