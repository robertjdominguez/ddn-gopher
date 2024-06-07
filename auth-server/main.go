package main

import (
	"dominguezdev.com/auth-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/healthcheck", handlers.HealthCheckHandler)
	router.POST("/login", handlers.LoginHandler)
	router.POST("/refresh", handlers.RefreshHandler)

	router.Run(":8080")
}
