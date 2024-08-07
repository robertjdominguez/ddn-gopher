package handlers

import (
	"fmt"
	"net/http"

	"dominguezdev.com/auth-server/models"
	"dominguezdev.com/auth-server/repository"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest models.User

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	foundUser, err := repository.CheckForUser(loginRequest.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}

	user, err := repository.VerifyUser(loginRequest.Password, foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is incorrect"})
		return
	}

	tokenString, err := repository.GenerateJWT(user.Username, user.ID, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
