package hasher

import (
	"crypto/sha256"
	"fmt"
	"os"
)

type IHasher interface {
	Hash(password string) string
}

type Hasher struct {
	salt string
}

func NewHasher() *Hasher {
	return &Hasher{
		salt: os.Getenv("SALT"),
	}
}

func (h *Hasher) Hash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
