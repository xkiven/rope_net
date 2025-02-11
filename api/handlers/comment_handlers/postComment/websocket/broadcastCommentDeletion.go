package websocket

func BroadcastCommentDeletion(postID uint, commentID uint) {
	PostClientsMutex.RLock()
	if clients, ok := PostClients[postID]; ok {
		for client := range clients {
			message := map[string]interface{}{
				"action":    "delete",
				"commentID": commentID,
			}
			err := client.WriteJSON(message)
			if err != nil {
				// 移除出错的客户端
				PostClientsMutex.Lock()
				delete(PostClients[postID], client)
				PostClientsMutex.Unlock()
			}
		}
	}
	PostClientsMutex.RUnlock()
}
