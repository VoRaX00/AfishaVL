package extension

import (
	"Afisha/internal/application"
)

type Services struct {
	application.IUserService
}

func NewServices(repos *Repositories) *Services {
	return &Services{
		IUserService: application.NewUserService(repos),
	}
}
