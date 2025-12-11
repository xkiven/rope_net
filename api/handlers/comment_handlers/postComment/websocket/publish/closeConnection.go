package publish

import (
	"github.com/gorilla/websocket"
)

// CloseConnection 关闭连接并从对应帖子的连接列表中移除客户端
func CloseConnection(ws *websocket.Conn, postID uint) {
	PostClientsMutex.Lock()
	if clients, ok := PostClients[postID]; ok {
		delete(clients, ws)
	}
	PostClientsMutex.Unlock()
	ws.Close()
}
