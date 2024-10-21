package extension

import (
	"Afisha/internal/application"
)

type Services struct {
	application.IUserService
	application.ITokenService
	application.IEventsService
}

func NewServices(repos *Repositories) *Services {
	return &Services{
		IUserService:   application.NewUserService(repos.IUsersRepository),
		ITokenService:  application.NewTokenService(repos.ITokenRepository),
		IEventsService: application.NewEventService(repos.IEventsRepository),
	}
}
