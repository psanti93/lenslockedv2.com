package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Password could not be hashed:%w", err)
	}

	hashedPassword := string(hashedByte)

	return hashedPassword, nil
}
