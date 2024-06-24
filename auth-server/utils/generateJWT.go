package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// This will generate a JWT with the username, id, and role.
func GenerateJWT(username string, id float64, roles ...string) (string, error) {
	// Let's make sure we have the secret
	jwtSecret := GetJWTSecret()

	// Determine the role based on the argument
	role := "admin"
	if len(roles) > 0 {
		role = roles[0]
	}

	claims := GenerateClaims(username, id, role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("GenerateJWT: failed to generate a JWT: %w", err)
	}

	return tokenString, nil
}
