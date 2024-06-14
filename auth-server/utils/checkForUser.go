package utils

import (
	"log"
)

func CheckForUser(username string) map[string]interface{} {
	client := CreateClient()

	// Define variables for the query
	variables := map[string]interface{}{
		"username": username,
	}

	token, err := GenerateJWT("", "admin")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the query
	respData, err := QueryHasura(client, UserQuery, variables, token)
	if err != nil {
		log.Fatal(err)
	}

	return respData
}
