package application

import "Afisha/internal/domain"

type IUserService interface {
	Create(user domain.UserRegister) (string, error)
	Update(user domain.User) error
	GetById(userId string) (domain.User, error)
	Verify(user domain.UserToLogin) (domain.User, error)
}
