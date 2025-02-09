package comment_handlers

import "github.com/gorilla/websocket"

// CloseConnection 关闭连接并从对应帖子的连接列表中移除客户端
func CloseConnection(ws *websocket.Conn, postID uint) {
	postClientsMutex.Lock()
	if clients, ok := postClients[postID]; ok {
		delete(clients, ws)
	}
	postClientsMutex.Unlock()
	ws.Close()
}
