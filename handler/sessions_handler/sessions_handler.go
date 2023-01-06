package sessions_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/derUbermenk/go-user_auth_api/service/session_service"
	"github.com/gin-gonic/gin"
)

func Create(ss session_service.SessionService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var new_session_request session_service.NewSessionRequest

		_ = c.ShouldBindJSON(&new_session_request)

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

		http.SetCookie(
			c.Writer,
			&http.Cookie{
				Name:  "session",
				Value: fmt.Sprintf("%v", id),
			},
		)

		c.JSON(
			http.StatusOK,
			nil,
		)
	}
}

func Delete(ss session_service.SessionService) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie("session")

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				nil,
			)

			return
		}

		cookie := &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		}

		http.SetCookie(
			c.Writer,
			cookie,
		)

		c.JSON(http.StatusOK, nil)
	}
}
