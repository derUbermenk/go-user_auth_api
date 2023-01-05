package user_service

type NewUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user interface{}, success bool, err error)
	FetchUserByEmail(email string) (user interface{}, err error)
	FetchUser(id int) (user interface{}, err error)
	FetchUserSelf(id int) (user interface{}, err error)
	DeleteUser(id int) (deleted_user_id int, err error)
}
