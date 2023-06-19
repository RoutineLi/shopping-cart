package types

import "github.com/gorilla/websocket"

type SendData struct {
	Client *websocket.Conn
	RecvId uint
}

type Session map[uint]*SendData //map<user_id, client>

var SessionMap = make(Session, 1024)
