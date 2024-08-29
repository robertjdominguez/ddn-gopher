package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type DecodedToken struct {
	Username string
	UserId   string
	IsValid  bool
}

var AdminClaims = map[string]interface{}{
	"x-hasura-role":         "admin",
	"x-hasura-default-role": "admin",
	"x-hasura-allowed-roles": []string{
		"admin",
	},
}

// ShapeUserClaims shapes user claims based on the userId and userRole.
func ShapeUserClaims(userId string, userRole string) map[string]interface{} {
	return map[string]interface{}{
		"x-hasura-user-id":       userId,
		"x-hasura-role":          userRole,
		"x-hasura-default-role":  "user",
		"x-hasura-allowed-roles": []string{"user", "admin"},
	}
}

// GenerateClaims generates JWT claims based on the username, userId, and userRole.
func GenerateClaims(username string, userId string, userRole string) jwt.MapClaims {
	claims := jwt.MapClaims{
		"user":                         username,
		"exp":                          time.Now().Add(time.Hour * 72).Unix(),
		"claims.jwt.hasura.io": map[string]interface{}{},
	}

	hasuraClaims := claims["claims.jwt.hasura.io"].(map[string]interface{})

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
