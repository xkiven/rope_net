package addPostComment

import (
	"Rope_Net/models"
	"Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gorilla/websocket"
)

// SendHistoricalComments 发送该帖子的历史评论给客户端
func SendHistoricalComments(ws *websocket.Conn, postID uint) error {
	dbConn, err := db.ConnectDB()
	if err != nil {
		logger.Error("连接数据库错误: %w", err)
		return err
	}
	defer db.CloseDB(dbConn)

	var comments []models.PostComment
	result := dbConn.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		logger.Error("查询历史评论失败: %w", result.Error)
		return result.Error
	}

	for _, comment := range comments {
		if err := ws.WriteJSON(comment); err != nil {
			logger.Error("发送历史评论给客户端失败: %w", err)
			return err
		}
	}
	return nil
}
