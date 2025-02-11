package publish

import (
	"Rope_Net/models"
)

func BroadcastComment(comment models.PostComment) {
	PostClientsMutex.RLock()
	if clients, ok := PostClients[comment.PostID]; ok {
		for client := range clients {
			err := client.WriteJSON(comment)
			if err != nil {
				// 移除出错的客户端
				PostClientsMutex.Lock()
				delete(PostClients[comment.PostID], client)
				PostClientsMutex.Unlock()
			}
		}
	}
	PostClientsMutex.RUnlock()
}
