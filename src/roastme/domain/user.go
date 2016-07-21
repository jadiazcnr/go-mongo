package user

import "time"

// User represents information about users.
type User struct {
	Name         string    `bson:"name"`
	Username     string    `bson:"username"`
	Password     string    `bson:"password"`
	Email        string    `bson:"email"`
	CreationDate time.Time `bson:"creationtime"`
}

// NewUser returns a new user
func NewUser(name, username, password, email string) User {
	return User{
		Name:         name,
		Username:     username,
		Password:     password,
		Email:        email,
		CreationDate: time.Now(),
	}
}
