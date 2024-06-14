package handlers

import (
	"net/http"
	"os"

	"dominguezdev.com/auth-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"
)

// TODO: Refactor this to match the login LoginHandler — we'll need utils for:
// 1. validating the token
// 2. validating the claims
// 3. Using the username in the claims instead of the raw string
func RefreshHandler(c *gin.Context) {
	var request struct {
		Token string `json:"token"`
	}

	// Load the .env file
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

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	username := claims["user"].(string)
	newTokenString, err := utils.GenerateJWT(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newTokenString})
}
