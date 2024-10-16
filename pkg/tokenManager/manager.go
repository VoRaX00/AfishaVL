package tokenManager

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserIp string `json:"user_ip"`
}

type ITokenManager interface {
	GetAccessToken(ipClient string, ttl time.Duration) (string, error)
	GetRefreshToken() (string, error)
	HashRefreshToken(token string) (string, error)
}

type TokenManager struct {
	signingKey string
}

func NewTokenManager(signingKey string) *TokenManager {
	return &TokenManager{
		signingKey: signingKey,
	}
}

func (t *TokenManager) GetAccessToken(ipClient string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserIp: ipClient,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *TokenManager) GetRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (t *TokenManager) HashRefreshToken(token string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(token))
	return hex.EncodeToString(hash.Sum(nil)), nil
}
