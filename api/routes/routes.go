package routes

import (
	"Rope_Net/api/handlers/comment_handlers/postComment/websocket/publish"
	"Rope_Net/api/handlers/comment_handlers/postComment/websocket/remove"
	"Rope_Net/api/handlers/comment_handlers/threadComment"
	"Rope_Net/api/handlers/post_handlers"
	"Rope_Net/api/handlers/user_handlers"
	"Rope_Net/api/handlers/user_handlers/login"
	"Rope_Net/middleware"
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
		post := r.Group("/post")
		//post.Use(middleware.IdentifyTokenMiddleware)
		{
			post.POST("/publish", middleware.IdentifyTokenMiddleware, post_handlers.PublishPost)
			post.GET("/getPost/:postID", post_handlers.GetPost)
			post.DELETE("/deletePost/:postID", middleware.IdentifyTokenMiddleware, post_handlers.DeletePost)
			post.GET("/getPostList", post_handlers.GetPostList)
		}
		comment := r.Group("/comment")
		{
			comment.GET("/ws", middleware.IdentifyTokenMiddleware, publish.WebSocketHandler)
			comment.DELETE("/deletePostComment/:postCommentID", middleware.IdentifyTokenMiddleware, remove.DeleteComment)
			comment.POST("/createThreadComment", middleware.IdentifyTokenMiddleware, threadComment.CreateThreadComment)
			comment.GET("/getThreadComment/:commentID", threadComment.GetThreadComment)
			comment.DELETE("/deleteThreadComment/:threadCommentID", middleware.IdentifyTokenMiddleware, threadComment.DeleteThreadComment)
		}
	}

}
