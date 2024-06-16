package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// This will generate a JWT with the username, id, and role.
func GenerateJWT(username string, roles ...string) (string, error) {
	// Check if JWT_SECRET is already set, otherwise load from .env file
	if os.Getenv("JWT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file and JWT_SECRET environment variable not set")
		}
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		panic("JWT_SECRET environment variable not set")
	}

	// Determine the role based on the argument
	role := "admin"
	if len(roles) > 0 {
		role = roles[0]
	}

	claims := GenerateClaims(username, 1, role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
