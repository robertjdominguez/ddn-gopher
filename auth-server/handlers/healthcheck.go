package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	// Errything is good

	c.JSON(http.StatusOK, gin.H{"message": "I'm alive you fools"})
}
