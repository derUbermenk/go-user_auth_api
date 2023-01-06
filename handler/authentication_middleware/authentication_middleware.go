package authentication_middleware

import "github.com/gin-gonic/gin"

// checks if client is logged in
func Authenticate() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// checks if the client owns the resource
func AuthorizeOwner() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
