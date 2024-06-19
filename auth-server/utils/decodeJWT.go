package utils

import (
	"errors"

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
