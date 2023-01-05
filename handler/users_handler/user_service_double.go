package users_handler

// Definition for the User Service Double used for the handler tests are defined here

import "github.com/derUbermenk/go-user_auth_api/service/user_service"

type UserServiceDouble struct{}

func (f *UserServiceDouble) CreateUser(newUserRequest user_service.NewUserRequest) (
	user interface{},
	success bool,
	err error) {

	if newUserRequest.Email == "invalid_email@email.com" {
		user = nil
		success = false
		err = nil
		return
	}

	user = nil
	success = true
	err = nil
	return
}

func (f *UserServiceDouble) FetchUserByEmail(email string) (user interface{}, err error) {
	if email == "existing_email@email.com" {
		user = struct{}{}
		return
	}
	user = nil
	return
}
