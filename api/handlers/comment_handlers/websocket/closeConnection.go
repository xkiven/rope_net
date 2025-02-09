package websocket

import (
	"Rope_Net/api/handlers/comment_handlers"
	"github.com/gorilla/websocket"
)

// CloseConnection 关闭连接并从对应帖子的连接列表中移除客户端
func CloseConnection(ws *websocket.Conn, postID uint) {
	comment_handlers.PostClientsMutex.Lock()
	if clients, ok := comment_handlers.PostClients[postID]; ok {
		delete(clients, ws)
	}
	comment_handlers.PostClientsMutex.Unlock()
	ws.Close()
}
