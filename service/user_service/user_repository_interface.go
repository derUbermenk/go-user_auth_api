package user_service

import "github.com/derUbermenk/go-user_auth_api/repository"

type UserRepositoryInterface interface {
	Create(user_info map[string]interface{}) (user repository.User, err error)
	FindByEmail(email string) (user repository.User, err error)
	FindPublic(id int) (user repository.User, err error)
	FindPrivate(id int) (user repository.User, err error)
	Delete(id int) (deleted_user repository.User, err error)
}

type UserRepositoryDouble struct{}

func (u *UserRepositoryDouble) Create(user_info map[string]interface{}) (user repository.User, err error) {
	return
}

func (u *UserRepositoryDouble) FindByEmail(email string) (user repository.User, err error) {
	if email == "existing_user_email@email.com" {
		user.ID = 1
		return
	}

	return
}

func (u *UserRepositoryDouble) FindPublic(id int) (user repository.User, err error) {
	if id == 1 {
		user.ID = 1
		return
	}

	return
}

func (u *UserRepositoryDouble) FindPrivate(id int) (user repository.User, err error) {
	if id == 1 {
		user.ID = 1
		return
	}

	return
}

func (u *UserRepositoryDouble) Delete(id int) (deleted_user repository.User, err error) {
	if id == 1 {
		deleted_user.ID = 1

		return
	}

	return
}
