package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const cost = 8

func HashPsw(psw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(psw), cost)
	if err != nil {
		return "", errors.New("failed to hash the psw")
	}

	return string(hashed), nil
}
func CheckPsw(hashedPsw, psw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPsw), []byte(psw))
}
