package authentication_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// checks if client is logged in
func Authenticate() func(c *gin.Context) {
	return func(c *gin.Context) {
		session_cookie, err := c.Request.Cookie("session")

		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", session_cookie.Value)
		c.Next()
	}
}

// checks if the client owns the resource
func AuthorizeOwner() func(c *gin.Context) {
	return func(c *gin.Context) {
		current_user_id, _ := c.Get("user_id")
		owner_id := c.Param("id")

		if current_user_id != owner_id {
			c.Status(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
