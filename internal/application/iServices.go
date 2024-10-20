package application

import "Afisha/internal/domain"

type IUserService interface {
	Create(user domain.UserRegister) (int, error)
	Update(user domain.User) error
	GetById(userId string) (domain.User, error)
	Verify(user domain.UserToLogin) (domain.User, error)
}

type ITokenService interface {
	AddRefreshToken(userId, ip, token string) error
	UpdateRefreshToken(userId, ip, token string) error
}

type IEventsService interface {
	Create(event domain.Event) (int, error)
	Update(event domain.Event) error
	GetById(eventId string) (domain.Event, error)
	Delete(eventId string) error
}
