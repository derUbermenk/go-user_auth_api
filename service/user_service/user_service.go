package user_service

type NewUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user interface{}, status bool, err error)
}
