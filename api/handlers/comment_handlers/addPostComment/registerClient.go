package addPostComment

import (
	"Rope_Net/api/handlers/comment_handlers"
	"github.com/gorilla/websocket"
)

// 注册客户端到对应帖子的连接列表
func registerClient(ws *websocket.Conn, postID uint) {
	comment_handlers.postClientsMutex.Lock()
	if comment_handlers.postClients[postID] == nil {
		comment_handlers.postClients[postID] = make(map[*websocket.Conn]bool)
	}
	comment_handlers.postClients[postID][ws] = true
	comment_handlers.postClientsMutex.Unlock()
}
