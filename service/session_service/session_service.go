package session_service

type SessionService interface {
	CreateSession(request NewSessionRequest) (id int, valid_credentials bool, err error)
}
