package comment_handlers

import "github.com/gorilla/websocket"

// 注册客户端到对应帖子的连接列表
func registerClient(ws *websocket.Conn, postID uint) {
	postClientsMutex.Lock()
	if postClients[postID] == nil {
		postClients[postID] = make(map[*websocket.Conn]bool)
	}
	postClients[postID][ws] = true
	postClientsMutex.Unlock()
}
