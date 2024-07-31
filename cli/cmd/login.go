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
	noPrompt bool
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with your username and password",
	Run: func(cmd *cobra.Command, args []string) {
		if noPrompt {
			err := Login(username, password)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	},
}

func init() {
	// This AddCommand adds the loginCmd to our root set of commands.
	// rootCmd is available because this file and it are part of the same
	// cmd package.
	rootCmd.AddCommand(loginCmd)

	// This sets the flags our command uses and marks which are req
	loginCmd.Flags().BoolVarP(&noPrompt, "no-prompt", "n", false, "Use command-line flags for login")
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}

// We're making this a public function so that we can reuse it in the TUI
func Login(username, password string) error {
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
		return fmt.Errorf("Error decoding response: %v", err)
	}

	// Let's first check to see if the server returns an error
	errorMsg, ok := result["error"].(string)
	if ok {
		return fmt.Errorf("Error logging in: %s", errorMsg)
	}

	// Otherwise, let's try to get the value from the "token" key of the returned data
	token, ok := result["token"].(string)
	if !ok {
		return fmt.Errorf("Error: token not found in response")
	}

	// For now, let's just print the token with an affirming message 🤷
	fmt.Println("You bastard...you logged in! Here's your token:", token)
	return nil
}
