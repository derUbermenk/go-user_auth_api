package sessions_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(
			http.StatusOK,
			map[string]string{
				"message": "hit context",
			},
		)
	}
}
