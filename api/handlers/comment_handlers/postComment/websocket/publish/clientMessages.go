package publish

import (
	"Rope_Net/models"
	"Rope_Net/pkg/logger"
	"github.com/gorilla/websocket"
	"time"
)

func HandleClientMessages(ws *websocket.Conn, postID uint, userID uint) {
	for {
		var comment models.PostComment
		err := ws.ReadJSON(&comment)
		if err != nil {
			logger.Error("读取客户端消息失败，帖子 ID: %d，错误信息: %v", postID, err)
			break
		}
		comment.PostID = postID
		comment.UserID = userID
		comment.CreateTime = time.Now()
		// 处理评论，比如保存到数据库
		if err := SaveCommentToDB(comment); err != nil {
			logger.Error("保存评论到数据库失败，帖子 ID: %d，评论内容: %v，错误信息: %v", postID, comment, err)
			continue
		}
		// 广播新的评论给订阅该帖子的客户端
		BroadcastComment(comment)
	}
}
