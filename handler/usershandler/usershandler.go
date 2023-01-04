package usershandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(
			http.StatusOK,
			map[string]string{
				"message": "hello gin",
			},
		)
	}
}

func Show() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		c.IndentedJSON(
			http.StatusOK,
			map[string]int{
				"id": id,
			},
		)
	}
}
