package models

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
	send chan []byte 
}
