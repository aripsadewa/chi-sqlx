package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const Bcrypt = 14

func GeneratePassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), Bcrypt)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
