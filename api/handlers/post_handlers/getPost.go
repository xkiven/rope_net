package post_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPost(c *gin.Context) {
	var ThePost models.Post
	//传入post的id
	logger.Info("绑定参数")
	if err := c.ShouldBindJSON(&ThePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"info":   "绑定参数错误",
		})
		logger.Error(err)
		return
	}
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
	result := db.Where("id = ?", ThePost.ID).First(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "此post不存在",
		})
		logger.Error(result.Error)
		return
	}
	//增加浏览量
	post.PageView++
	//修改数据
	logger.Info("增加浏览量并存入数据库")
	result = db.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "修改帖子失败",
		})
		logger.Error(result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"data":   post,
	})
}
