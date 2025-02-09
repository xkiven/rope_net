package comment_handlers

func BroadcastCommentDeletion(postID uint, commentID uint) {
	postClientsMutex.RLock()
	if clients, ok := postClients[postID]; ok {
		for client := range clients {
			message := map[string]interface{}{
				"action":    "delete",
				"commentID": commentID,
			}
			err := client.WriteJSON(message)
			if err != nil {
				// 移除出错的客户端
				postClientsMutex.Lock()
				delete(postClients[postID], client)
				postClientsMutex.Unlock()
			}
		}
	}
	postClientsMutex.RUnlock()
}
