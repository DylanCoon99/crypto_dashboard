package controllers

import (
	"log"
	"sync"
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Represents a connected ws client
type Client struct {
	ClientID uuid.UUID
	Conn     *websocket.Conn
}

type Manager struct {
	Clients map[uuid.UUID]*Client
	Send    chan []byte
	mu      sync.Mutex
}

var manager = &Manager{
	Clients: make(map[uuid.UUID]*Client),
	Send:    make(chan []byte, 256),
}

func (cfg *ApiConfig) HandleWebSocket(c *gin.Context) {

	ws, err := cfg.Upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Websocket upgrader error:", err)
		return
	}

	defer ws.Close()

	// Generate unique UUID for the client
	clientUUID := uuid.New()
	client := &Client{ClientID: clientUUID, Conn: ws}

	// Client is added to manager
	manager.mu.Lock()
	manager.Clients[clientUUID] = client
	manager.mu.Unlock()

	log.Printf("Client %s has successfully connected via websocket. Total clients: %d", clientUUID, len(manager.Clients))

	// defer disconnect of client; delete the client from the manager map
	defer func() {
		manager.mu.Lock()
		delete(manager.Clients, clientUUID)
		ws.Close()
		manager.mu.Unlock()
		log.Printf("Client %s has successfully disconnected. Total clients: %d", clientUUID, len(manager.Clients))
	}()

	// if we want to listen to Client messages, we would do so here


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 
	go client.writeMessages(ctx)


	<-ctx.Done()

}



func (c *Client) writeMessages(ctx context.Context) {

	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}

}


func Broadcast(ctx context.Context) {

	for message	:= range manager.Send {
		for _, client := range manager.Clients {

			err := client.Conn.WriteMessage(websocket.TextMessage, message)

			if err != nil {
				log.Printf("Websocket write error: %v", err)
				return
			}
			log.Printf("Message sent to client %s: %s", client.ClientID.String(), string(message))

		}
	}


}


func GetRealTimePrices(ctx context.Context) {

	// gets realtime time price data every 5 mins and writes to the managers Send channel

	for {
		manager.mu.Lock()
		time.Sleep(2 * time.Second)

		log.Println("Simulating api call...")
		manager.Send <- []byte("test")
		manager.mu.Unlock()
	}

}
