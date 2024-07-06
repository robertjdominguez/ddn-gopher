package utils

import (
	"errors"
	"log"
)

func CheckForUser(username string) (map[string]interface{}, error) {
	client := CreateClient()

	// Check to see if there is a non-empty string for the username
	if username == "" {
		return nil, errors.New("username cannot be an empty string")
	}

	// Define variables for the query
	variables := map[string]interface{}{
		"username": username,
	}

	token, err := GenerateJWT("", 0, "admin")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the query
	respData, err := QueryHasura(client, UserQuery, variables, token)
	if err != nil {
		log.Fatal(err)
	}

	return respData, nil
}
