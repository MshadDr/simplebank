package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hashpassword return a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword( []byte(password), bcrypt.DefaultCost )
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string( hashedPassword ), nil
}

// CheckPassword check if the provided password is match with its hash
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword( []byte(hashedPassword), []byte(password) )
}
