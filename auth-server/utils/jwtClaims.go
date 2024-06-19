package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var AdminClaims = map[string]interface{}{
	"x-hasura-role":         "admin",
	"x-hasura-default-role": "admin",
	"x-hasura-allowed-roles": []string{
		"admin",
	},
}

// Function to shape user claims
func ShapeUserClaims(userId float64, userRole string) map[string]interface{} {
	return map[string]interface{}{
		"x-hasura-user-id":       userId,
		"x-hasura-role":          userRole,
		"x-hasura-default-role":  "user",
		"x-hasura-allowed-roles": []string{"user", "admin"},
	}
}

// Function to generate the claims
func GenerateClaims(username string, userId float64, userRole string) jwt.MapClaims {
	claims := jwt.MapClaims{
		"user":                         username,
		"exp":                          time.Now().Add(time.Hour * 72).Unix(),
		"https://hasura.io/jwt/claims": map[string]interface{}{},
	}

	hasuraClaims := claims["https://hasura.io/jwt/claims"].(map[string]interface{})

	if userRole == "admin" {
		for key, value := range AdminClaims {
			hasuraClaims[fmt.Sprint(key)] = value
		}
	} else {
		userClaims := ShapeUserClaims(userId, userRole)
		for key, value := range userClaims {
			hasuraClaims[fmt.Sprint(key)] = value
		}
	}

	return claims
}
