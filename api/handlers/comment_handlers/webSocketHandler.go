package comment_handlers

import (
	"Rope_Net/api/handlers/comment_handlers/addPostComment"
	"Rope_Net/models"
	"Rope_Net/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// 定义 websocket 升级器
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 存储每个帖子 ID 对应的 WebSocket 连接
var postClients = make(map[uint]map[*websocket.Conn]bool)
var postClientsMutex sync.RWMutex

func WebSocketHandler(c *gin.Context) {
	postIDStr := c.Query("postID")
	var postID uint
	_, err := fmt.Sscanf(postIDStr, "%d", &postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"info":   "无效的帖子 ID",
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
	// 升级 HTTP 连接为 WebSocket 连接
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("WebSocket 升级失败:", err)
		return
	}
	defer addPostComment.CloseConnection(ws, postID)

	// 注册客户端到对应帖子的连接列表
	addPostComment.registerClient(ws, postID)

	// 发送该帖子的历史评论给客户端
	if err := addPostComment.SendHistoricalComments(ws, postID); err != nil {
		logger.Error("发送历史评论失败，帖子 ID: %d，错误信息: %v", postID, err)
		return
	}

	// 监听客户端消息
	addPostComment.HandleClientMessages(ws, postID, userID)

}
