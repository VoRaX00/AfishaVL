package infrastructure

import (
	"Afisha/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: *db,
	}
}

func (r *UserRepository) Create(user domain.User) (string, error) {
	return "", nil
}

func (r *UserRepository) Update(user domain.User) error {
	return nil
}

func (r *UserRepository) GetById(userId string) (domain.User, error) {
	return domain.User{}, nil
}
