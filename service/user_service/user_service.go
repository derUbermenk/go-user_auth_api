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
}

type UserServ struct {
	user_db UserRepositoryInterface
}

func (u *UserServ) CreateUser(newUserRequest NewUserRequest) (user interface{}, success bool, err error) {
	// check presence of required fields
	valid := newUserRequest.is_valid()

	if !valid {
		success = false
		return
	}

	user_info := newUserRequest.to_map()
	user, err = u.user_db.Create(user_info)

	return
}
