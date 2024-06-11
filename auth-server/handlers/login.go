package handlers

import (
	"encoding/json"
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
    query MyQuery($username: User_Varchar) {
      user_users(where: {username: {_eq: $username}}) {
        id
        username
        password
      }
    }
	`

	// Define variables for the query
	variables := map[string]interface{}{
		"username": credentials.Username,
	}

	token, err := GenerateJWT("", "admin")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the query
	respData, err := graphqlClient.QueryHasura(client, query, variables, token)
	if err != nil {
		log.Fatal(err)
	}

	// Let's find out about the user
	user, err := json.Marshal(respData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(string(user))

	// If everything is good, let's give them a token with a user role
	tokenString, err := GenerateJWT(credentials.Username, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
