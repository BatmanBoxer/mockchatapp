package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handlers) AddClient(userID string, conn *websocket.Conn) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	client := &Client{
		Conn:    conn,
		Message: make(chan string),
	}
	(*h.conn)[userID] = client

	go h.handleMessages(client)
}

func (h *Handlers) RemoveClient(userID string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	client, ok := (*h.conn)[userID]
	if !ok {
		return
	}

	close(client.Message)

	delete(*h.conn, userID)
}



func (h *Handlers) handleMessages(client *Client) {
	for message := range client.Message {
		err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error sending message to client", client.Message, err)
			return
		}
	}
}

func (h *Handlers)handleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	h.AddClient(userID, conn)

	if err != nil {
		fmt.Println("Error sending message:", err)
  }
  //make a broadcast msg function and client disscont
}









