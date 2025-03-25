package handlers

import (
	"fmt"
	"net/http"
	"time"
	"github.com/batmanboxer/mockchatapp/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handlers) addClient(chatRoomId string, client *models.Client) {
	h.mutex.Lock()
	h.mutex.Unlock()

	h.client[chatRoomId] = append((h.client[chatRoomId]), client)

	go h.handleMessages(client)
	//go h.testMsg(client)
  go h.listenMessage(chatRoomId,client)
}

func (h *Handlers) removeClient(chatRoomId string, userId string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	clients, ok := h.client[chatRoomId]
	if !ok {
		return
	}

	var updatedClients []*models.Client
	for _, client := range clients {
		if client.Id != userId {
			updatedClients = append(updatedClients, client)
		} else {
			if client.Messagech != nil {
				close(client.Messagech)
			}
		}
	}

	if len(updatedClients) == 0 {
		delete(h.client, chatRoomId)
	} else {
		h.client[chatRoomId] = updatedClients
	}
}

func (h *Handlers) listenMessage(roomID string, client *models.Client) {
	//authorized client
	for {
		messageType, p, err := client.Conn.ReadMessage()
		if err != nil {
			client.Closech <- struct{}{}
			break
		}
		if messageType != websocket.TextMessage {
			continue
		}
		// message := models.Message{}
		// err = json.Unmarshal(p,&message)
		// if err != nil{
		//   client.Closech<-struct{}{}
		//   break
		// }

		h.broadcastClient(roomID, string(p))
	}

}

func (h *Handlers) broadcastClient(roomId string, message string) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	clients, ok := h.client[roomId]
	if !ok {
		return
	}

	for _, client := range clients {
		if client.Messagech != nil {
			client.Messagech <- message
		}
	}
}

func (h *Handlers) testMsg(client *models.Client) {
	for {
		client.Messagech <- "testing"
		time.Sleep(5 * time.Second)
	}
}

func (h *Handlers) handleMessages(client *models.Client) {
	for message := range client.Messagech {
		err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error sending message to client", client.Messagech, err)
			return
		}
	}
}
func (h *Handlers) Listenhandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	clientId := vars["id"]
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return err
	}
	client := &models.Client{
		Id:        clientId,
		Conn:      conn,
		Messagech: make(chan string),
		Closech:   make(chan struct{}),
	}

	h.addClient("testChatRoom", client)

	<-client.Closech
	return nil
}
