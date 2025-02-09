package post_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostList(c *gin.Context) {
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
	//获得数据
	logger.Info("获取PostList")
	var post []models.Post
	result := db.Order("page_view DESC").Find(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "读取数据失败",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"data":   post,
	})

}
