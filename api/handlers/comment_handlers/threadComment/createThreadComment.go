package threadComment

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateThreadComment(c *gin.Context) {
	//绑定参数
	var threadComment models.ThreadComments
	if err := c.ShouldBindJSON(&threadComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10001,
			"info":   "绑定参数失败",
		})
		logger.Info(err)
		return
	}
	// 从上下文中获取用户信息
	logger.Info("从上下文获取用户信息")
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10011,
			"info":   "无法获取用户信息"})
		return
	}
	currentUser := user.(*models.User)
	userID := currentUser.Id
	//连接数据库
	db, err := db2.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10002,
			"error":  "连接数据库错误",
		})
		logger.Error(err)
		return
	}
	defer db2.CloseDB(db)
	//插入嵌套评论
	threadComment.UserID = userID
	threadComment.CreateTime = time.Now()
	result := db.Create(&threadComment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "发布嵌套评论失败",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "发布嵌套评论成功",
	})

}
