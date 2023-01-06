package session_service

type NewSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionService interface {
	CreateSession(request NewSessionRequest) (id int, valid_credentials bool, err error)
}
