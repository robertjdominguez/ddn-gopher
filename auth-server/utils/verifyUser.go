package utils

import (
	"fmt"
)

// VerifyUser checks if the provided password matches any user in the userInformation map.
// Returns a boolean indicating success and a message with the result.
func VerifyUser(reqPassword string, userInformation map[string]interface{}) (bool, string) {
	if userUsers, ok := userInformation["user_users"].([]interface{}); ok {
		if len(userUsers) == 0 {
			return false, "No users found"
		}
		for _, user := range userUsers {
			if userMap, ok := user.(map[string]interface{}); ok {
				if id, ok := userMap["id"].(float64); ok {
					fmt.Println("User ID:", int(id))
				}
				if password, ok := userMap["password"].(string); ok {
					if password == reqPassword {
						println("Password is correct. Here's a token!")
						return true, "Password is correct"
					} else {
						return false, "Password incorrect"
					}
				} else {
					return false, "Password not found for user"
				}
			} else {
				return false, "User data is not in expected format"
			}
		}
	} else {
		return false, "User users data not found or is in incorrect format"
	}
	return false, "User not found"
}
