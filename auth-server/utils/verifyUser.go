package utils

import (
	"errors"
)

// VerifyUser checks if the provided password matches any user in the userInformation map.
// Returns a boolean indicating success and a message with the result.
func VerifyUser(reqPassword string, userInformation map[string]interface{}) (string, float64, error) {
	if userUsers, ok := userInformation["user_users"].([]interface{}); ok {
		if len(userUsers) == 0 {
			return "", 0, errors.New("no users found")
		}
		for _, user := range userUsers {
			if userMap, ok := user.(map[string]interface{}); ok {
				if id, ok := userMap["id"].(float64); ok {
					if password, ok := userMap["password"].(string); ok {
						if password == reqPassword {
							if username, ok := userMap["username"].(string); ok {
								return username, id, nil
							}
							return "", 0, errors.New("username not found for user")
						} else {
							return "", 0, errors.New("password is incorrect")
						}
					} else {
						return "", 0, errors.New("password not found for user")
					}
				} else {
					return "", 0, errors.New("user ID not found or is in incorrect format")
				}
			} else {
				return "", 0, errors.New("user data is not in expected format")
			}
		}
	} else {
		return "", 0, errors.New("user data not found or is in incorrect format")
	}
	return "", 0, errors.New("user not found")
}
