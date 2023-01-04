package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	v1_routes(r)

	r.Run(":8080")
}

func v1_routes(r *gin.Engine) {

	/*
		// session related routes
		sessions := r.Group("/sessions")
		{
			sessions.POST("", sessionshandler.Create())
			sessions.DELETE("", app_middlewares.RequireAuth(), sessions_handler.Destroy(SessionService))
		}

		// user related routes
		r.POST("/users/create", usershandler.Create())

		users_protected := r.Group("/users")
		users_protected.Use(app_middlewares.RequireAuth())
		{
			users_protected.GET("/:id", usershandler.Show())
			users_protected.DELETE("/:id/destroy", userHandler.Destroy(UserService))
			users_protected.PUT("/:id/update", userHandler.Update(UserService))
		}
	*/
}
