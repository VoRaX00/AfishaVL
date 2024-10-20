package extension

import (
	"Afisha/internal/application"
	"Afisha/internal/infrastructure"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	application.IUsersRepository
	application.ITokenRepository
	application.IEventsRepository
	db *sqlx.DB
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		IUsersRepository:  infrastructure.NewUserRepository(db),
		ITokenRepository:  infrastructure.NewTokenRepository(db),
		IEventsRepository: infrastructure.NewEventRepository(db),
		db:                db,
	}
}
