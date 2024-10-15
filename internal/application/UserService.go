package application

import "Afisha/internal/domain"

type UserService struct {
	repo IUsersRepository
}

func NewUserService(repo IUsersRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user domain.User) (string, error) {
	return s.repo.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) GetById(userId string) (domain.User, error) {
	return s.repo.GetById(userId)
}
