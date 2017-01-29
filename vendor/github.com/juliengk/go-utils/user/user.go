package user

import (
	"os/user"
)

type User struct {
	*user.User
}

func New() (*User, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &User{user}, nil
}

func (u *User) IsRoot() bool {
	if u.Uid == "0" {
		return true
	}

	return false
}
