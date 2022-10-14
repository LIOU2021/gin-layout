package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashCheck(hash string, key string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(key))
	if err != nil {
		return false
	} else {
		return true
	}

}
