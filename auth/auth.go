package auth

import (
	"fmt"
)

type user struct {
	Name     string
	Username string
	Password string
	Sessions []string
}

var User user;

func LookupUser(cookie string) (user, error) {
	for _, s := range User.Sessions {
		if s == cookie{
			return User, nil
		}
	}
	return user{}, fmt.Errorf("no user with given session token")
}
