package extension

import (
	"Afisha/internal/application"
	"Afisha/internal/infrastructure"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	application.IUsersRepository
	db *sqlx.DB
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		IUsersRepository: infrastructure.NewUserRepository(db),
		db:               db,
	}
}
