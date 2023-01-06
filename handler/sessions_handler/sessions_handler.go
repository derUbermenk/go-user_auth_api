package sessions_handler

import (
	"log"
	"net/http"

	"github.com/derUbermenk/go-user_auth_api/service/session_service"
	"github.com/gin-gonic/gin"
)

func Create(ss session_service.SessionService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var new_session_request session_service.NewSessionRequest

		_ = c.ShouldBindJSON(session_service.new_session_request)

		id, valid_credentials, err := ss.CreateSession(new_session_request)

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				nil,
			)

			return
		}

		if !valid_credentials {
			c.JSON(
				http.StatusUnauthorized,
				nil,
			)

			return
		}

		c.SetCookie("session", id)

		c.JSON(
			http.StatusOK,
			nil,
		)
	}
}
