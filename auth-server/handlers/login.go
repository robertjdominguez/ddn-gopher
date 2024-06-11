package handlers

import (
	"fmt"
	"log"
	"net/http"

	"dominguezdev.com/auth-server/graphqlClient"
	"github.com/gin-gonic/gin"
)

/* To log a user in, we'll use a credentials struct and then we'll generate
* a token for them.
 */

var credentials struct {
	Username string `json:"username"`
	Passowrd string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	// If the request isn't properly formatted JSON using credentials, error
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: Check that this is a real user
	// We'll need to hit the GraphQL endpoint (which should be `3000`) via an env var and query if the username exists.
	// If it does, we'll check the password; if it doesn't we'll create a new user.
	// If the password is correct, we'll send them a token as a response; if they're new,
	// they'll get one, too.
	client := graphqlClient.CreateClient()

	// Here's our GraphQL query to check if a user exists
	query := `
		query($username: string!) {
	     user_users {
	       id
	       username
	     }
		}
	`

	// Define variables for the query
	variables := map[string]interface{}{
		"username": credentials.Username,
	}

	// Execute the query
	respData, err := graphqlClient.QueryHasura(client, query, variables)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response
	user := respData["user"].(map[string]interface{})
	fmt.Printf("User: %s\n", user["username"])

	// If everything is good, let's give them a token
	tokenString, err := GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
