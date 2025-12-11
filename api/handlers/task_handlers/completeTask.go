package task_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CompleteTask(c *gin.Context) {
	logger.Info("绑定参数")
	taskID := c.Param("taskID")
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
	//查找此任务
	var task models.Task
	result := db.Where("id = ? AND user_id = ?", taskID, currentUser.Id).Find(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "查找数据库失败",
		})
		logger.Error(result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"info":   "不存在此任务",
		})
		return
	}
	//修改数据
	task.Completed = true
	//插入数据库
	result = db.Save(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "修改任务失败",
		})
		logger.Error(result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 10006,
		"info":   "已完成任务",
	})
}
