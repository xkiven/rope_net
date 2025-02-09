package addPostComment

import (
	"Rope_Net/api/handlers/comment_handlers"
	"Rope_Net/models"
)

func BroadcastComment(comment models.PostComment) {
	comment_handlers.postClientsMutex.RLock()
	if clients, ok := comment_handlers.postClients[comment.PostID]; ok {
		for client := range clients {
			err := client.WriteJSON(comment)
			if err != nil {
				// 移除出错的客户端
				comment_handlers.postClientsMutex.Lock()
				delete(comment_handlers.postClients[comment.PostID], client)
				comment_handlers.postClientsMutex.Unlock()
			}
		}
	}
	comment_handlers.postClientsMutex.RUnlock()
}
