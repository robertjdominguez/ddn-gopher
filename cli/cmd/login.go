package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// Top-level variables for our username and password being passed by the user
var (
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with your username and password",
	Run: func(cmd *cobra.Command, args []string) {
		login(username, password)
	},
}

func init() {
	// This AddCommand adds the loginCmd to our root set of commands
	// TODO: Where is our rootCmd coming from in this context?
	rootCmd.AddCommand(loginCmd)

	// This sets the flags our command uses and marks which are req
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}

func login(username, password string) {
	// We'll cretae a payload that will eventually be transformed into JSON
	payload := map[string]string{
		"username": username,
		"password": password,
	}

	// We'll attempt to transform it to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error creating JSON payload for username and password:", err)
	}

	// Then, we'll try sending it along...
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	// We'll defer a closing of the connection, too
	defer resp.Body.Close()

	// We'll make sure we can decode the JSON returned by the auth server
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
	}

	// Let's first check to see if the server returns an error
	error, ok := result["error"].(string)

	// If there is, we need to do an early return and let the user know what the error is
	if ok {
		fmt.Println("Error logging in:", error)
		return
	}

	// Otherwise, let's try to get the value from the "token" key of the returned data
	token, ok := result["token"].(string)

	// On the chance "token" is missing, something bad happened
	if !ok {
		fmt.Println("Error: token not found in response")
		return
	}

	// For now, let's just print the token with an affirming message 🤷
	fmt.Println("You bastard...you logged in! Here's your token:", token)
}
