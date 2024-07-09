package models

import "errors"

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	ID       float64 `json:"id"`
}

// VerifyPassword can be used on found user to check if their povided password is correct.
func (u *User) VerifyPassword(reqPassword string) error {
	if reqPassword != u.Password {
		return errors.New("password is incorrect")
	}
	return nil
}
