package handlers

import (
	"net/http"

	"dominguezdev.com/auth-server/models"
	"dominguezdev.com/auth-server/utils"
	"github.com/gin-gonic/gin"
)

func RefreshHandler(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// If there is no error, we have a token we can decode
	decodedToken, err := utils.DecodeJWT(request.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newTokenString, err := utils.GenerateJWT(decodedToken.Username, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newTokenString})
}
