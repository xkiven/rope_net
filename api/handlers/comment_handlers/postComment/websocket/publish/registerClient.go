package publish

import (
	"github.com/gorilla/websocket"
)

// 注册客户端到对应帖子的连接列表
func RegisterClient(ws *websocket.Conn, postID uint) {
	PostClientsMutex.Lock()
	if PostClients[postID] == nil {
		PostClients[postID] = make(map[*websocket.Conn]bool)
	}
	PostClients[postID][ws] = true
	PostClientsMutex.Unlock()
}
