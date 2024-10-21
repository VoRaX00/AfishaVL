package application

import (
	"Afisha/internal/domain"
)

type UserService struct {
	userRepo IUsersRepository
}

func NewUserService(userRepo IUsersRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(user domain.UserRegister) (int, error) {
	return s.userRepo.Create(user)
}

func (s *UserService) GetById(userId int) (domain.User, error) {
	return s.userRepo.GetById(userId)
}

func (s *UserService) Verify(user domain.UserToLogin) (domain.User, error) {
	return s.userRepo.Verify(user)
}
