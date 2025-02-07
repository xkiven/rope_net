package routes

import (
	"Rope_Net/api/handlers/user_handlers"
	"Rope_Net/api/handlers/user_handlers/login"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Group("/api")
	{
		user := r.Group("/user")
		{
			user.POST("/register", user_handlers.Register)
			user.POST("/PreLogin", login.PreLogin)
		}
	}

}
