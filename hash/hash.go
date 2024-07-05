package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func HashCheck(data string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err != nil
}
