package websocket

import "Rope_Net/api/handlers/comment_handlers"

func BroadcastCommentDeletion(postID uint, commentID uint) {
	comment_handlers.PostClientsMutex.RLock()
	if clients, ok := comment_handlers.PostClients[postID]; ok {
		for client := range clients {
			message := map[string]interface{}{
				"action":    "delete",
				"commentID": commentID,
			}
			err := client.WriteJSON(message)
			if err != nil {
				// 移除出错的客户端
				comment_handlers.PostClientsMutex.Lock()
				delete(comment_handlers.PostClients[postID], client)
				comment_handlers.PostClientsMutex.Unlock()
			}
		}
	}
	comment_handlers.PostClientsMutex.RUnlock()
}
