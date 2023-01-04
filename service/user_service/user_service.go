package user_service

type NewUserRequest struct{}

type UserService interface {
	CreateUser(newUserRequest NewUserRequest) (user user_data.User, status bool, err error)
}
