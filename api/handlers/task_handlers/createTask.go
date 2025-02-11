package task_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateTask(c *gin.Context) {
	var task models.Task
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
	var newTask struct {
		Name     string `json:"name"`
		Deadline string `json:"deadline"`
	}
	logger.Info("绑定参数")
	if err := c.ShouldBind(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10001,
			"info":   "绑定参数失败",
		})
		logger.Info(err)
		return
	}
	deadline, err := time.Parse(time.RFC3339, newTask.Deadline)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10013,
			"error":  "无效的截止时间格式，使用 RFC3339 格式 (YYYY-MM-DDTHH:MM:SSZ)",
		})
		logger.Error(err)
		return
	}
	//创建任务
	task.UserID = currentUser.Id
	task.Deadline = deadline
	task.Name = newTask.Name
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
	//插入数据库
	logger.Info("创建任务并存入数据库")
	result := db.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "创建任务失败",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "创建任务成功",
	})

}
