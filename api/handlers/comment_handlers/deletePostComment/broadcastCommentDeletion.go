package deletePostComment

import "Rope_Net/api/handlers/comment_handlers"

func BroadcastCommentDeletion(postID uint, commentID uint) {
	comment_handlers.postClientsMutex.RLock()
	if clients, ok := comment_handlers.postClients[postID]; ok {
		for client := range clients {
			message := map[string]interface{}{
				"action":    "delete",
				"commentID": commentID,
			}
			err := client.WriteJSON(message)
			if err != nil {
				// 移除出错的客户端
				comment_handlers.postClientsMutex.Lock()
				delete(comment_handlers.postClients[postID], client)
				comment_handlers.postClientsMutex.Unlock()
			}
		}
	}
	comment_handlers.postClientsMutex.RUnlock()
}
