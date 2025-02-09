package websocket

import (
	"Rope_Net/api/handlers/comment_handlers"
	"github.com/gorilla/websocket"
)

// 注册客户端到对应帖子的连接列表
func RegisterClient(ws *websocket.Conn, postID uint) {
	comment_handlers.PostClientsMutex.Lock()
	if comment_handlers.PostClients[postID] == nil {
		comment_handlers.PostClients[postID] = make(map[*websocket.Conn]bool)
	}
	comment_handlers.PostClients[postID][ws] = true
	comment_handlers.PostClientsMutex.Unlock()
}
