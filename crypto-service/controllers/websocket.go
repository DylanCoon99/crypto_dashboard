package controllers

import (
	"log"
	"sync"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Represents a connected ws client
type Client struct {
	ClientID uuid.UUID
	Conn     *websocket.Conn
	Send     chan []byte
}

type Manager struct {
	Clients map[uuid.UUID]*Client
	mu      sync.Mutex
}

var manager = &Manager{
	Clients: make(map[uuid.UUID]*Client),
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
	client := &Client{ClientID: clientUUID, Conn: ws, Send: make(chan byte, 256)}

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
		log.Printf("Client %s has stopped writing messages.", c.ClientID.String())
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


func (m *Manager) Broadcast(message string) {

}
