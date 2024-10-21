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

func (r *UserRepository) GetById(userId int) (domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Verify(user domain.UserToLogin) (domain.User, error) {
	var usr domain.User
	query := "SELECT * FROM users WHERE phone = $1 AND password_hash = $2"
	row := r.db.QueryRow(query, usr.Login, user.Password)
	if err := row.Scan(&usr); err != nil {
		return domain.User{}, err
	}
	return usr, nil
}
