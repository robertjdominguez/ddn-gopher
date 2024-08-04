package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"dominguezdev.com/auth-server/models"
	"github.com/golang-jwt/jwt/v4"
)

func DecodeJWT(encodedToken string) (models.DecodedToken, error) {
	var decodedToken models.DecodedToken

	// Let's make sure we have the secret
	jwtSecret := GetJWTSecret()

	// Then, we can parse the JWT
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// If we get an error, there's something wrong with the token
	if err != nil || !token.Valid {
		return decodedToken, errors.New("invalid token")
	}

	// We'll see if the claims are accurate — if they aren't, or the token isn't valid
	// then we'll throw an error
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return decodedToken, errors.New("invalid token")
	}

	// We'll check to make sure our Hasura claims are part of the JWT
	hasuraClaims, ok := claims["https://hasura.io/jwt/claims"].(map[string]interface{})
	if !ok {
		return decodedToken, errors.New("invalid Hasura claims")
	}

	// Then, we'll make sure there's a user present
	userId, ok := hasuraClaims["x-hasura-user-id"].(float64)
	if !ok {
		return decodedToken, errors.New("user ID not found in claims")
	}

	// Finally, we'll send back the decoded information
	decodedToken = models.DecodedToken{
		IsValid:  true,
		UserId:   userId,
		Username: claims["user"].(string),
	}

	return decodedToken, nil
}

// GetJWTSecret loads the JWT_SECRET from the environment variables or .env file
func GetJWTSecret() []byte {
	// Load the .env file if JWT_SECRET is not already set
	if os.Getenv("JWT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file and JWT_SECRET environment variable not set: %v", err)
		}
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	return jwtSecret
}

// This will generate a JWT with the username, id, and role.
func GenerateJWT(username string, id float64, roles ...string) (*string, error) {
	// Let's make sure we have the secret
	jwtSecret := GetJWTSecret()

	// Determine the role based on the argument
	role := "admin"
	if len(roles) > 0 {
		role = roles[0]
	}

	claims := models.GenerateClaims(username, id, role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("GenerateJWT: failed to generate a JWT: %w", err)
	}

	return &tokenString, nil
}
