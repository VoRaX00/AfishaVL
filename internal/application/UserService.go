package application

import (
	"Afisha/internal/domain"
)

type UserService struct {
	userRepo  IUsersRepository
	tokenRepo ITokenRepository
}

func NewUserService(userRepo IUsersRepository, tokenRepo ITokenRepository) *UserService {
	return &UserService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *UserService) Create(user domain.UserRegister) (int, error) {
	return s.userRepo.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) GetById(userId string) (domain.User, error) {
	return s.userRepo.GetById(userId)
}

func (s *UserService) Verify(user domain.UserToLogin) (domain.User, error) {
	return s.userRepo.Verify(user)
}

func (s *UserService) AddRefreshToken(userId, ip, token string) error {
	return s.tokenRepo.Create(userId, ip, token)
}

func (s *UserService) UpdateRefreshToken(userId, ip, token string) error {
	return s.tokenRepo.Update(userId, ip, token)
}
