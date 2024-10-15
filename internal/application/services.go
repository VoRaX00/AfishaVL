package application

import "Afisha/internal/infrastructure"

type Services struct {
	IUserService
}

func NewServices(repos *infrastructure.Repositories) *Services {
	return &Services{
		IUserService: NewUserService(repos),
	}
}
