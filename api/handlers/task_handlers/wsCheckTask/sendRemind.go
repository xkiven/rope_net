package wsCheckTask

import (
	"Rope_Net/pkg/logger"
	"github.com/gorilla/websocket"
)

func SendRemind(userID uint, message string) {
	for _, client := range userClients[userID] {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			logger.Info("发送消息到用户 %d 的客户端失败: %v", userID, err)
			client.Close()
			// 移除该连接
			for i, c := range userClients[userID] {
				if c == client {
					userClients[userID] = append(userClients[userID][:i], userClients[userID][i+1:]...)
					break
				}
			}
		}
	}
}
