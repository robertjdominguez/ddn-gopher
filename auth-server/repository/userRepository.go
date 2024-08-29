package repository

import (
	"strconv"

	"dominguezdev.com/auth-server/errors"
	"dominguezdev.com/auth-server/models"
	"dominguezdev.com/auth-server/utils"
)

func CheckForUser(username string) (models.User, error) {
	client := utils.CreateClient()

	if username == "" {
		return models.User{}, customErrors.ErrEmptyUsername
	}

	variables := map[string]interface{}{
		"username": username,
	}

	// We'll need an admin-level JWT to check for the presence of the user
	token, err := GenerateJWT("", "", "admin")
	if err != nil {
		return models.User{}, err
	}

	respData, err := utils.QueryHasura(client, utils.UserQuery, variables, *token)
	if err != nil {
		return models.User{}, err
	}

	returnedUser, ok := respData["usersByUsername"]
	if !ok {
		return models.User{}, customErrors.ErrUserDataNotFound
	}

	userMap, ok := returnedUser.(map[string]interface{})
	if !ok {
		return models.User{}, customErrors.ErrUserDataFormat
	}

	id, ok := userMap["id"].(float64)
	if !ok {
		return models.User{}, customErrors.ErrUserIDFormat
	}

	idStr := strconv.FormatFloat(id, 'f', -1, 64)

	username, ok = userMap["username"].(string)
	if !ok {
		return models.User{}, customErrors.ErrUsernameNotFound
	}

	password, ok := userMap["password"].(string)
	if !ok {
		return models.User{}, customErrors.ErrPasswordNotFound
	}

	user := models.User{
		Username: username,
		Password: password,
		ID:       idStr,
	}

	return user, nil
}

func VerifyUser(reqPassword string, user models.User) (*models.User, error) {
	if err := user.VerifyPassword(reqPassword); err == nil {
		return &user, nil
	}
	return nil, customErrors.ErrInvalidUsernameOrPassword
}
