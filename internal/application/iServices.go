package application

import "Afisha/internal/domain"

type IUserService interface {
	Create(user domain.User) (string, error)
	Update(user domain.User) error
	GetById(userId string) (domain.User, error)
}
