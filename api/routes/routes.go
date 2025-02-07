package routes

import (
	"Rope_Net/api/handlers/user_handlers"
	"Rope_Net/api/handlers/user_handlers/login"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(sessions.Sessions("mySession", sessions.NewCookieStore([]byte("secret"))))
	r.Group("/api")
	{
		user := r.Group("/user")
		{
			user.POST("/register", user_handlers.Register)
			user.POST("/preLogin", login.PreLogin)
			user.POST("/finalLogin", login.FinalLogin)
		}
	}

}
