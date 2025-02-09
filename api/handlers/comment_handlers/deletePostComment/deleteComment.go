package deletePostComment

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteComment(c *gin.Context) {
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
	userID := currentUser.Id
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

	//验证是否为评论作者本人
	logger.Info("验证是否为评论作者本人")
	var postComment models.PostComment
	result := db.Where("id = ?", commentID).First(&postComment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "查找失败",
		})
		logger.Error(result.Error)
		return

	}
	if postComment.UserID != userID {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"info":   "你不是此评论的主人",
		})
		return
	}

	//数据库中删除
	logger.Info("在数据库中删除")
	result = db.Delete(&models.PostComment{}, postComment.ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "删除失败",
		})
		logger.Error(result.Error)
		return
	}

	BroadcastCommentDeletion(postComment.PostID, commentID)

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "删除评论成功",
	})
}
