package wsCheckTask

import (
	"Rope_Net/models"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 升级 HTTP 连接为 WebSocket 连接
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 存储每个用户 ID 对应的 WebSocket 连接
var userClients = make(map[uint][]*websocket.Conn)

func WsHandler(c *gin.Context) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	defer conn.Close()

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
	userID := currentUser.Id
	// 将新连接添加到对应用户的客户端列表中
	userClients[userID] = append(userClients[userID], conn)

	// 保持连接
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			// 移除该连接
			for i, client := range userClients[userID] {
				if client == conn {
					userClients[userID] = append(userClients[userID][:i], userClients[userID][i+1:]...)
					break
				}
			}
			break
		}
	}

}
