package utils

import (
	"fmt"
)

func VerifyUser(reqPassword string, userInformation map[string]interface{}) bool {
	if userUsers, ok := userInformation["user_users"].([]interface{}); ok {
		for _, user := range userUsers {
			// Here, we'll create the userMap and assert the type is, essentially, JSON
			if userMap, ok := user.(map[string]interface{}); ok {
				// We're doing the same thing here with the id and asserting it's a float
				if id, ok := userMap["id"].(float64); ok {
					fmt.Println("User ID:", int(id))
				}
				// Then again with the password to do a check
				if password, ok := userMap["password"].(string); ok {
					if password == reqPassword {
						println("Password is correct. Here's a token!")
						return true
					} else {
						return false
					}
				}
			} else {
				fmt.Println("User data is not in expected format")
				return false
			}
		}
	} else {
		fmt.Println("User users data not found or is in incorrect format")
		return false
	}
	return false
}
