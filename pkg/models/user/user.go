package user

import (
)

// User is a model of user
type User struct {
	ID			int64	`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
	Role 		string	`json:"role"`
}

func copyUser(u User) *User {
	return &User{
		ID: u.ID, 
		Username: u.Username, 
		Password: u.Password, 
		Role: u.Role,
	}
}

// NewUser Return a new User
func NewUser(username, password, role string) *User {
	return &User{
			Username: username,
			Password: password,
			Role: role,
		}
}