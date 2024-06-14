package utils

import (
	"os"
	"time"

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
	// BUG: This only works with an x-hausra-user-id incldued...even with admin
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-user-id":       "1",
			"x-hasura-role":          role,
			"x-hasura-default-role":  "user",
			"x-hasura-allowed-roles": []string{"user", "admin"},
		},
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
