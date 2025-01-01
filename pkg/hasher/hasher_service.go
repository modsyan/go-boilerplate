package hasher

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password using bcrypt.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares a plain-text password with a hashed password.
func CheckPasswordHash(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
