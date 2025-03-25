package models

import (
	"github.com/gorilla/websocket"
	"sync"
)

// remove sender later and check it via jwt
type Message struct {
	Message string
}

type Client struct {
	Id        string
	Conn      *websocket.Conn
	Messagech chan string
	Closech   chan struct{}
	Mutex     *sync.RWMutex
}
