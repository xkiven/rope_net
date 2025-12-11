package post_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PublishPost(c *gin.Context) {
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

	//解析帖子，绑定参数
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10001,
			"info":   "绑定参数失败",
		})
		logger.Info(err)
		return
	}
	//设置帖子的用户ID
	post.UserID = currentUser.Id
	post.PublishTime = time.Now()

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

	//插入数据
	logger.Info("创建帖子并存入数据库")
	result := db.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "发布帖子失败",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "发布成功",
	})

}
