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

func (r *UserRepository) Create(user domain.UserRegister) (string, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user)
}

func (r *UserRepository) Update(user domain.User) error {
	return nil
}

func (r *UserRepository) GetById(userId string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepository) Verify(user domain.UserToLogin) (domain.User, error) {

}
