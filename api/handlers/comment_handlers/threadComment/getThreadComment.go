package threadComment

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetThreadComment(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	var commentID uint
	_, err := fmt.Sscanf(commentIDStr, "%d", &commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"info":   "无效的评论 ID",
		})
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
	//查找此评论的嵌套评论
	logger.Info("查找此评论的嵌套评论")
	var threadComments []models.ThreadComments
	result := db.Where("comment_id = ?", commentID).Find(&threadComments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "查找数据库错误",
		})
		logger.Error(result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"data":   threadComments,
	})

}
