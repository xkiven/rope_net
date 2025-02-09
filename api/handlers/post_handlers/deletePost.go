package post_handlers

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeletePost(c *gin.Context) {
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

	//传入要删的post的id
	logger.Info("绑定参数")
	postID := c.Param("postID")

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

	//验证身份是否为post的创建者
	logger.Info("检查此post是否存在")
	var post models.Post
	result := db.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "此post不存在",
		})
		logger.Error(result.Error)
		return
	}
	if post.UserID != currentUser.Id {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10013,
			"info":   "你不是此post作者",
		})
		return
	}

	//删除post
	result = db.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"info":   "删除post失败",
		})
		logger.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "你已成功删除此post",
	})

}
