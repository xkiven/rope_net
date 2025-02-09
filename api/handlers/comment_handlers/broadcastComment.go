package comment_handlers

import "Rope_Net/models"

func BroadcastComment(comment models.PostComment) {
	postClientsMutex.RLock()
	if clients, ok := postClients[comment.PostID]; ok {
		for client := range clients {
			err := client.WriteJSON(comment)
			if err != nil {
				// 移除出错的客户端
				postClientsMutex.Lock()
				delete(postClients[comment.PostID], client)
				postClientsMutex.Unlock()
			}
		}
	}
	postClientsMutex.RUnlock()
}
