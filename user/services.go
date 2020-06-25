package main

import (
	"github.com/google/uuid"
)

var userService *UserService

type UserService struct{}

func (s *UserService) CreateUser(user *User) error {
	query := "INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3)"

	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	if _, err = db.Exec(query, uuidV4, user.Email, user.PasswordHash); err != nil {
		return err
	}

	return nil
}
