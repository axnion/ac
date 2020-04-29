package users

import "github.com/google/uuid"

type User struct {
	ID       string `jsonapi:"primary,user"`
	Username string `jsonapi:"attr,username"`
	Password string `jsonapi:"attr,password"`
}

var savedUsers []*User

func GetUsers() []*User {
	return savedUsers
}

func GetUsersWithId() []*User {
	return savedUsers
}

func Write(user *User) {
	if user.ID == "" {
		id, _ := uuid.NewRandom()
		user.ID = id.String()
	}

	savedUsers = append(savedUsers, user)
}
