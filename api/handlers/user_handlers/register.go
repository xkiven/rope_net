package user_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	logger.Info("绑定参数")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"error":  "绑定参数错误",
		})
		logger.Error(err)
		return
	}
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
	logger.Info("检查用户名是否已存在")
	var existingUser models.User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10003,
			"error":  "用户名已存在",
		})
		logger.Error(result.Error)
		return
	} else if result.Error.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"error":  "查询数据库错误",
		})
		logger.Error(result.Error)
		return
	}

	logger.Info("插入数据库")
	result = db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"error":  "插入数据库错误",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "success",
	})
}
