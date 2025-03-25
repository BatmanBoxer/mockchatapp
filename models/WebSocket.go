package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
  Id string
	Conn *websocket.Conn
  Messagech chan string
  Closech chan struct{}
  Mutex *sync.RWMutex
}
