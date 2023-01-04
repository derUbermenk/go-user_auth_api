package users_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(user_service api.UserService) func(c *gin.Context) {
	type CreateResponse struct {
		User_ID int `json:user_id`
	}

	return func(c *gin.Context) {
		var new_user_request api.NewUserRequest

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

		user, status, err := user_service.CreateUser(new_user_request)

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				nil,
			)

			return
		}

		if status == false {
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
