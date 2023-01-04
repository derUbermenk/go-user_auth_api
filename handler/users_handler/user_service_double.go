package users_handler

// Definition for the User Service Double used for the handler tests are defined here

import "github.com/derUbermenk/go-user_auth_api/service/user_service"

type UserServiceDouble struct{}

func (f *UserServiceDouble) CreateUser(newUserRequest user_service.NewUserRequest) (
	user interface{},
	success bool,
	err error) {

	if newUserRequest.Email == "invalid_email@email.com" {
		success = false
		return
	}

	return user, success, err
}
