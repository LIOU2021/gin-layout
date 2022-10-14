package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashMake(key string) string {
	hash1, _ := bcrypt.GenerateFromPassword([]byte(key), bcrypt.MinCost)
	return string(hash1)
}
