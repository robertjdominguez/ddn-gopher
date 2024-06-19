package handlers

import (
	"fmt"
	"net/http"

	"dominguezdev.com/auth-server/models"
	"dominguezdev.com/auth-server/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var credentials models.Credentials

	// If the request isn't properly formatted JSON using credentials, error
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Otherwise, we'll take the username and see if it exists
	respData := utils.CheckForUser(credentials.Username)

	// Then, if we've found them, we'll verify their credentials
	isVerified, err := utils.VerifyUser(credentials.Password, respData)

	// If everything is good, let's give them a token with a user role
	if !isVerified {
		fmt.Printf("Error: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// If everything is good, let's give them a token with a user role
	tokenString, err := utils.GenerateJWT(credentials.Username, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Then, serve it up to them
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
