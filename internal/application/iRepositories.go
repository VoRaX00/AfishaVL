package application

import (
	"Afisha/internal/domain"
	"time"
)

type IUsersRepository interface {
	Create(user domain.UserRegister) (int, error)
	GetById(userId int) (domain.User, error)
	Verify(user domain.UserToLogin) (domain.User, error)
}

type ITokenRepository interface {
	Create(token domain.RefreshToken) error
	Update(newTokenHash, tokenHash, ipClient string, ttl time.Duration) error
}

type IEventsRepository interface {
	Create(event domain.Event) (int, error)
	Update(event domain.Event) error
	GetById(eventId string) (domain.Event, error)
	Delete(eventId string) error
}
