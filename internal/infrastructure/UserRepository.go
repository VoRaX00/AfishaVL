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

func (r *UserRepository) Create(user domain.UserRegister) (int, error) {
	var id int
	query := "INSERT INTO users (phone, password_hash) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, user.Phone, user.Password, &id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) Update(user domain.User) error {
	return nil
}

func (r *UserRepository) GetById(userId string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepository) Verify(user domain.UserToLogin) (domain.User, error) {
	return domain.User{}, nil
}
