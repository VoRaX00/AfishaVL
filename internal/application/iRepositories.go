package application

import "Afisha/internal/domain"

type IUsersRepository interface {
	Create(user domain.UserRegister) (int, error)
	Update(user domain.User) error
	GetById(userId string) (domain.User, error)
	Verify(user domain.UserToLogin) (domain.User, error)
}

type ITokenRepository interface {
	Create(userId, ip, token string) error
	Update(userId, ip, token string) error
}

type IEventsRepository interface {
	Create(event domain.Event) (int, error)
	Update(event domain.Event) error
	GetById(eventId string) (domain.Event, error)
	Delete(eventId string) error
}
