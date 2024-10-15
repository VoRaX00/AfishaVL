package infrastructure

import (
	"Afisha/internal/application"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	application.IUsersRepository
	db *sqlx.DB
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		IUsersRepository: NewUserRepository(db),
		db:               db,
	}
}
