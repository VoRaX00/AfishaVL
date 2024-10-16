package extension

import (
	"Afisha/internal/application"
)

type Services struct {
	application.IUserService
	application.ITokenRepository
}

func NewServices(repos *Repositories) *Services {
	return &Services{
		IUserService: application.NewUserService(repos.IUsersRepository, repos.ITokenRepository),
	}
}
