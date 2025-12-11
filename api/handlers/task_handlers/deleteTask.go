package task_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTask(c *gin.Context) {
	// 从上下文中获取用户信息
	logger.Info("从上下文获取用户信息")
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10011,
			"info":   "无法获取用户信息",
		})
		return
	}
	currentUser := user.(*models.User)
	//传入要删掉的taskID
	logger.Info("绑定参数")
	taskID := c.Param("taskID")
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

	//检查此任务
	logger.Info("检查此任务")
	var task models.Task
	result := db.Where("id = ?", taskID).First(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "查找数据库错误",
		})
		logger.Error(result.Error.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"info":   "没找到此任务",
		})
		return
	}
	if task.UserID != currentUser.Id {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10013,
			"info":   "你不是此任务的主人",
		})
		return
	}
	if task.Completed == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10016,
			"info":   "你的任务还未完成无法删除",
		})
		return
	}

	//删除任务
	result = db.Delete(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "删除任务时失败",
		})
		logger.Error(result.Error.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "删除任务成功",
	})

}
