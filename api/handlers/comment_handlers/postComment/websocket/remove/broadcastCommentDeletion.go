package remove

import "Rope_Net/api/handlers/comment_handlers/postComment/websocket/publish"

func BroadcastCommentDeletion(postID uint, commentID uint) {
	publish.PostClientsMutex.RLock()
	if clients, ok := publish.PostClients[postID]; ok {
		for client := range clients {
			message := map[string]interface{}{
				"action":    "remove",
				"commentID": commentID,
			}
			err := client.WriteJSON(message)
			if err != nil {
				// 移除出错的客户端
				publish.PostClientsMutex.Lock()
				delete(publish.PostClients[postID], client)
				publish.PostClientsMutex.Unlock()
			}
		}
	}
	publish.PostClientsMutex.RUnlock()
}
