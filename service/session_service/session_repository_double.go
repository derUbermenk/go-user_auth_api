package session_service

import "github.com/derUbermenk/go-user_auth_api/repository"

type SessionRepositoryInterface interface {
	FindByEmail(email string) (user repository.User, err error)
}

type SessionRepositoryDouble struct{}

func (u SessionRepositoryDouble) FindByEmail(email string) (user repository.User, err error) {
	if email == "existing_email@email.com" {
		user.ID = 1
		user.Email = email
		user.Password = "valid_password"
		return
	}

	return
}
