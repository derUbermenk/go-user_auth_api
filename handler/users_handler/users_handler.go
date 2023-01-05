package users_handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/derUbermenk/go-user_auth_api/service/user_service"
	"github.com/gin-gonic/gin"
)

func Create(us user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var new_user_request user_service.NewUserRequest

		_ = c.ShouldBindJSON(&new_user_request)

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

func FetchByEmail(us user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")

		user, err := us.FetchUserByEmail(email)

		if err != nil {
			log.Printf("Service Error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				nil,
			)

			return
		}

		if user == nil {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func Fetch(us user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		requested_user_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		user, err := us.FetchUser(requested_user_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func FetchSelf(us user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		requested_user_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		user, err := us.FetchUserSelf(requested_user_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
