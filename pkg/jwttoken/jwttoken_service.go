package jwttoken

import (
	"company-name/configs"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

// GenerateToken generates a JWT token for user verification
func GenerateToken(userID string, expirationTime time.Time) string {
	config := configs.GetConfig()
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		log.Fatalf("Error generating JWT token: %v", err)
	}

	return signedToken
}

func GenerateAccessToken(userID string) string {
	config := configs.GetConfig()

	expirationTime := time.Now().Add(time.Duration(config.JWT.Expiration) * time.Millisecond)

	payload := map[string]interface{}{
		"userId": userID,
		"type":   "access",
	}

	tokenContent := jwt.MapClaims{
		"sub":     userID,
		"payload": payload,
		"exp":     jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	signedToken, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		log.Fatalf("Error generating JWT token: %v", err)
	}

	return signedToken
}

func ValidateAccessToken(token string) (jwt.MapClaims, error) {
	config := configs.GetConfig()
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	return claims, nil
}
