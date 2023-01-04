package users_handler

import (
	"log"
	"net/http"

	"github.com/derUbermenk/go-user_auth_api/service/user_service"
	"github.com/gin-gonic/gin"
)

func Create(us user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var new_user_request user_service.NewUserRequest

		_ = c.ShouldBindJSON(&new_user_request)

		/*
			if err != nil {
				log.Printf("Handler Error: %v", err)
				c.JSON(
					http.StatusBadRequest,
					nil,
				)
			}
		*/

		user, success, err := us.CreateUser(new_user_request)

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				nil,
			)

			return
		}

		if !success {
			c.JSON(
				http.StatusBadRequest,
				nil,
			)

			return
		}

		c.JSON(
			http.StatusOK,
			user,
		)
	}

}
