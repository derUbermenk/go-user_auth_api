package session_service

import "github.com/derUbermenk/go-user_auth_api/repository"

type NewSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionService interface {
	CreateSession(request NewSessionRequest) (id int, is_valid_credentials bool, err error)
}

func NewSessionService(db SessionRepositoryInterface) SessionService {
	return &sessionservice{
		db: db,
	}
}

type sessionservice struct {
	db SessionRepositoryInterface
}

func (s *sessionservice) CreateSession(request NewSessionRequest) (id int, is_valid_credentials bool, err error) {
	// validate credentials
	user, err := s.db.FindByEmail(request.Email)

	// no user found
	if user.ID == 0 {
		return
	}

	// do something about it
	if request.Password != user.Password {
		user = repository.User{}
		return
	}

	id = user.ID
	is_valid_credentials = true

	return
}
