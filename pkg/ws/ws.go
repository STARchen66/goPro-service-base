package ws

import "github.com/gorilla/websocket"

var upgrade = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    0,
	WriteBufferSize:   0,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}
