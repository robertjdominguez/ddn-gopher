package utils

import (
	"errors"
	"fmt"
)

// VerifyUser checks if the provided password matches any user in the userInformation map.
// Returns a boolean indicating success and a message with the result.
func VerifyUser(reqPassword string, userInformation map[string]interface{}) (bool, error) {
	if userUsers, ok := userInformation["user_users"].([]interface{}); ok {
		if len(userUsers) == 0 {
			return false, errors.New("no users found")
		}
		for _, user := range userUsers {
			if userMap, ok := user.(map[string]interface{}); ok {
				if id, ok := userMap["id"].(float64); ok {
					fmt.Println("User ID:", int(id))
				}
				if password, ok := userMap["password"].(string); ok {
					if password == reqPassword {
						println("Password is correct. Here's a token!")
						return true, nil
					} else {
						return false, errors.New("password is incorrect")
					}
				} else {
					return false, errors.New("pasword not found for user")
				}
			} else {
				return false, errors.New("user data is not in expected format")
			}
		}
	} else {
		return false, errors.New("user data not found or is in incorrect format")
	}
	return false, errors.New("user not found")
}
