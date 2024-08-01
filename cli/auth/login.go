package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// We're making this a public function so that we can reuse it in both the traditional
// CLI and the TUI.
func Login(username, password string) (string, error) {
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
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	// We'll defer a closing of the connection, too
	defer resp.Body.Close()

	// We'll make sure we can decode the JSON returned by the auth server
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("Error decoding response: %v", err)
	}

	// Let's first check to see if the server returns an error
	errorMsg, ok := result["error"].(string)
	if ok {
		return "", fmt.Errorf("Error logging in: %s", errorMsg)
	}

	// Otherwise, let's try to get the value from the "token" key of the returned data
	token, ok := result["token"].(string)
	if !ok {
		return "", fmt.Errorf("Error: token not found in response")
	}

	// For now, let's just print the token with an affirming message 🤷
	fmt.Println("You bastard...you logged in! Here's your token:", token)
	return token, nil
}
