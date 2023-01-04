package user_service

type NewUserRequest struct{}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user interface{}, status bool, err error)
}
