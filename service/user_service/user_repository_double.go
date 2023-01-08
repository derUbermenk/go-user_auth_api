package user_service

import "github.com/derUbermenk/go-user_auth_api/repository/user_repository"

type UserRepositoryDouble struct{}

func (u *UserRepositoryDouble) Create(user_info map[string]interface{}) (user user_repository.User, err error) {
	return
}

func (u *UserRepositoryDouble) FindByEmail(email string) (user user_repository.User, err error) {
	if email == "existing_user_email@email.com" {
		user.ID = 1
		return
	}

	return
}

func (u *UserRepositoryDouble) FindPublic(id int) (user user_repository.User, err error) {
	if id == 1 {
		user.ID = 1
		return
	}

	return
}

func (u *UserRepositoryDouble) FindPrivate(id int) (user user_repository.User, err error) {
	if id == 1 {
		user.ID = 1
		return
	}

	return
}
