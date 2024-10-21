package application

import "Afisha/internal/domain"

type IUserService interface {
	Create(user domain.UserRegister) (int, error)
	GetById(userId int) (domain.User, error)
	Verify(user domain.UserToLogin) (domain.User, error)
}

type ITokenService interface {
	AddToken(userId int, ip, token string) error
	GenerateTokens(ip string) (map[string]string, error)
	RefreshTokens(token, ip string) (map[string]string, error)
}

type IEventsService interface {
	Create(event domain.CreateEvent) (int, error)
	Update(event domain.Event) error
	GetById(eventId string) (domain.Event, error)
	Delete(eventId string) error
}
