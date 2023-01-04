package user_service

import "github.com/derUbermenk/go-user_auth_api/repository"

type NewUserRequest struct{}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user repository.User, status bool, err error)
}
