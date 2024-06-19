package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
