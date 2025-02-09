package comment_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateAndPublishComment(c *gin.Context) {
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
	//绑定参数
	logger.Info("绑定参数")
	var postComment models.PostComment
	if err := c.ShouldBindJSON(&postComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"info":   "绑定参数错误",
		})
		logger.Error(err.Error())
		return
	}
	postComment.UserID = currentUser.Id
	postComment.CreateTime = time.Now()
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
	//查找此post
	logger.Info("检查此post是否存在")
	var post models.Post
	result := db.Where("id = ?", postComment.PostID).First(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "此post不存在",
		})
		logger.Error(result.Error)
		return
	}
	//插入数据
	logger.Info("创建评论并存入数据库")
	result = db.Create(&postComment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "创建评论失败",
		})
		logger.Error(result.Error)
		return
	}
	//发布评论
	err = PublishComment(postComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10006,
			"info":   "发布评论失败",
		})
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "创建并发布评论成功",
	})

}
