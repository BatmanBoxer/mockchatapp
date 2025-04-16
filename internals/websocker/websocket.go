package websocker

import (
	"sync"
	"github.com/batmanboxer/mockchatapp/models"
)

type WebSocketManager struct {
	Client map[string][]*models.Client
	Mutex  *sync.RWMutex
}


