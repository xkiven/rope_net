package websocket

import (
	"Rope_Net/api/handlers/comment_handlers"
	"Rope_Net/models"
)

func BroadcastComment(comment models.PostComment) {
	comment_handlers.PostClientsMutex.RLock()
	if clients, ok := comment_handlers.PostClients[comment.PostID]; ok {
		for client := range clients {
			err := client.WriteJSON(comment)
			if err != nil {
				// 移除出错的客户端
				comment_handlers.PostClientsMutex.Lock()
				delete(comment_handlers.PostClients[comment.PostID], client)
				comment_handlers.PostClientsMutex.Unlock()
			}
		}
	}
	comment_handlers.PostClientsMutex.RUnlock()
}
