package sessions_handler

import "github.com/derUbermenk/go-user_auth_api/service/session_service"

type SessionServiceDouble struct{}

func (f *SessionServiceDouble) CreateSession(new_session_request session_service.NewSessionRequest) (id int, valid_credentials bool, err error) {
	if new_session_request.Password == "valid_password@email.com" {
		id = 1
		valid_credentials = true
		return
	}

	return
}
