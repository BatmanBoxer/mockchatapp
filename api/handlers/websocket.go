package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/batmanboxer/mockchatapp/common"
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
	h.webSocketManager.Mutex.Lock()
	h.webSocketManager.Mutex.Unlock()

	h.webSocketManager.Client[chatRoomId] = append(( h.webSocketManager.Client[chatRoomId]), client)

	go h.handleMessages(client)
	//go h.testMsg(client)
	go h.listenMessage(chatRoomId, client)
}

func (h *Handlers) initialMessage(chatRoomId string, client *models.Client, limit int) {
	messages, err := h.webSocketManager.GetMessages(chatRoomId, limit, 0)
	if err != nil {
		return
	}
	for _, message := range messages {
		client.Messagech <- message.Message
	}
}

func (h *Handlers) removeClient(chatRoomId string, userId string) {
	h.webSocketManager.Mutex.Lock()
	defer h.webSocketManager.Mutex.Unlock()

	clients, ok := h.webSocketManager.Client[chatRoomId]
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
		delete (h.webSocketManager.Client, chatRoomId)
	} else {
		 h.webSocketManager.Client[chatRoomId] = updatedClients
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

		h.broadcastMessage(roomID, string(p), client)
	}

}

func (h *Handlers) broadcastMessage(roomId string, message string, client *models.Client) {
	 h.webSocketManager.Mutex.RLock()
	defer h.webSocketManager.Mutex.RUnlock()

	clients, ok :=  h.webSocketManager.Client[roomId]
	if !ok {
		return
	}
	err := h.websocketManager.AddMessage(models.MessageModel{
		RoomId:   roomId,
		Message:  message,
		SenderId: client.Id,
	})
	if err != nil {
		log.Println(err.Error())
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

func (h *Handlers) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	chatroomId := vars["id"]
	userId := r.Context().Value(common.CONTEXTIDKEY)
	stringUserId := userId.(string)

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
		Id:        stringUserId,
		Conn:      conn,
		Messagech: make(chan string),
		Closech:   make(chan struct{}),
	}

	h.addClient(chatroomId, client)

	<-client.Closech
	conn.Close()
	h.removeClient(chatroomId, stringUserId)
	return nil
}
