package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Hash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Error while computing hash: %w", err)
	}
	return string(hash), nil
}

func HashMatchesString(hash, text string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text)) == nil
}
