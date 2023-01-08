package user_service

import (
	"encoding/json"
	"net/mail"

	"github.com/derUbermenk/go-user_auth_api/repository/user_repository"
)

type NewUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (req *NewUserRequest) is_valid() (valid bool) {
	// assert email is valid
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return
	}

	if req.Name == "" {
		return
	}

	if req.Password == "" {
		return
	}

	return true
}

func (req *NewUserRequest) to_map() (info map[string]interface{}) {
	data, _ := json.Marshal(req)

	json.Unmarshal(data, &info)

	return
}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user interface{}, success bool, err error)
	FetchUserByEmail(email string) (user interface{}, err error)
	FetchUser(id int) (user interface{}, err error)
	FetchUserSelf(id int) (user interface{}, err error)
	DeleteUser(id int) (deleted_user interface{}, err error)
}

type UserRepositoryInterface interface {
	Create(user_info map[string]interface{}) (user user_repository.User, err error)
	FindByEmail(email string) (user user_repository.User, err error)
	FindPublic(id int) (user user_repository.User, err error)
	FindPrivate(id int) (user user_repository.User, err error)
}

type userservice struct {
	user_db UserRepositoryInterface
}

func NewUserService(db UserRepositoryInterface) UserService {
	return &userservice{
		user_db: db,
	}
}

func (u *userservice) CreateUser(newUserRequest NewUserRequest) (user interface{}, success bool, err error) {
	// check presence of required fields
	valid := newUserRequest.is_valid()

	if !valid {
		success = false
		return
	}

	user_info := newUserRequest.to_map()
	user, err = u.user_db.Create(user_info)
	success = true
	return
}

func (u *userservice) FetchUserByEmail(email string) (user interface{}, err error) {
	fetched_user, err := u.user_db.FindByEmail(email)

	if fetched_user.ID == 0 {
		return
	}

	user = fetched_user
	return
}

func (u *userservice) FetchUser(id int) (user interface{}, err error) {
	public_user, err := u.user_db.FindPublic(id)

	if public_user.ID == 0 {
		return
	}

	user = public_user
	return
}

func (u *userservice) FetchUserSelf(id int) (user interface{}, err error) {
	private, err := u.user_db.FindPublic(id)

	if private.ID == 0 {
		return
	}

	user = private
	return
}

func (u *userservice) DeleteUser(id int) (deleted_user interface{}, err error) {
	return
}
