package repository

import (
	"errors"

	"dominguezdev.com/auth-server/models"
	"dominguezdev.com/auth-server/utils"
)

func CheckForUser(username string) ([]models.User, error) {
	client := utils.CreateClient()

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	variables := map[string]interface{}{
		"username": username,
	}

	token, err := GenerateJWT("", 0, "admin")
	if err != nil {
		return nil, err
	}

	respData, err := utils.QueryHasura(client, utils.UserQuery, variables, token)
	if err != nil {
		return nil, err
	}

	userUsers, ok := respData["user_users"].([]interface{})
	if !ok {
		return nil, errors.New("user data not found or is in incorrect format")
	}

	if len(userUsers) == 0 {
		return nil, errors.New("no users found")
	}

	var users []models.User
	for _, user := range userUsers {
		userMap, ok := user.(map[string]interface{})
		if !ok {
			return nil, errors.New("user data is not in expected format")
		}

		id, ok := userMap["id"].(float64)
		if !ok {
			return nil, errors.New("user ID not found or is in incorrect format")
		}

		username, ok := userMap["username"].(string)
		if !ok {
			return nil, errors.New("username not found for user")
		}

		password, ok := userMap["password"].(string)
		if !ok {
			return nil, errors.New("password not found for user")
		}

		users = append(users, models.User{
			Username: username,
			Password: password,
			ID:       id,
		})
	}

	return users, nil
}

func VerifyUser(reqPassword string, users []models.User) (*models.User, error) {
	for _, user := range users {
		if err := user.VerifyPassword(reqPassword); err == nil {
			return &user, nil
		}
	}
	return nil, errors.New("invalid username or password")
}
