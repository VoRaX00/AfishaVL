package application

import (
	"Afisha/internal/domain"
	"Afisha/pkg/tokenManager"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	accessTokenTTL  = time.Hour * 24
	refreshTokenTTL = time.Hour * 48
)

type TokenService struct {
	repo    ITokenRepository
	manager tokenManager.ITokenManager
}

func NewTokenService(repo ITokenRepository) *TokenService {
	signingKey := os.Getenv("SIGNING_KEY")

	return &TokenService{
		repo:    repo,
		manager: tokenManager.NewTokenManager(signingKey),
	}
}

func (s *TokenService) AddToken(userId int, ip, refreshToken string) error {
	hashToken, err := s.manager.HashRefreshToken(refreshToken)
	if err != nil {
		logrus.Error(err)
		return err
	}

	token := domain.RefreshToken{
		UserId:           userId,
		Ip:               ip,
		RefreshTokenHash: hashToken,
		ExpiresAt:        time.Now().Add(refreshTokenTTL),
	}

	err = s.repo.Create(token)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return s.repo.Create(token)
}

func (s *TokenService) GenerateTokens(ip string) (map[string]string, error) {
	tokens := make(map[string]string)
	accessToken, err := s.manager.GetAccessToken(ip, accessTokenTTL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	tokens["accessToken"] = accessToken
	refreshToken, err := s.manager.GetRefreshToken()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	tokens["refreshToken"] = refreshToken
	return tokens, nil
}

func (s *TokenService) RefreshTokens(token, ip string) (map[string]string, error) {
	tokens := make(map[string]string)

	refreshToken, err := s.manager.GetRefreshToken()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	hashNewToken, err := s.manager.HashRefreshToken(refreshToken)
	if err != nil {
		logrus.Error(err)
	}

	hashToken, err := s.manager.HashRefreshToken(token)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	err = s.repo.Update(hashNewToken, hashToken, ip, refreshTokenTTL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	accessToken, err := s.manager.GetAccessToken(ip, accessTokenTTL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	tokens["accessToken"] = accessToken
	tokens["refreshToken"] = refreshToken
	return tokens, nil
}
