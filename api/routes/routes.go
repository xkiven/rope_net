package routes

import (
	"Rope_Net/api/handlers/user_handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Group("/api")
	{
		user := r.Group("/user")
		{
			user.POST("/register", user_handlers.Register)
		}
	}

}
