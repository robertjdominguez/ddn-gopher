package customErrors

import "errors"

var (
	ErrEmptyUsername             = errors.New("username cannot be empty")
	ErrUserDataNotFound          = errors.New("user data not found or is in incorrect format")
	ErrNoUsersFound              = errors.New("no users found")
	ErrUserDataFormat            = errors.New("user data is not in expected format")
	ErrUserIDFormat              = errors.New("user ID not found or is in incorrect format")
	ErrUsernameNotFound          = errors.New("username not found for user")
	ErrPasswordNotFound          = errors.New("password not found for user")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
)
