package controllers

import (
	"context"
	"log"
	"time"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/DylanCoon99/crypto_dashboard/crypto-service/api"

)

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
	client := &Client{ClientID: clientUUID, Conn: ws, Send: make(chan []byte, 256)}

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a done channel to signal this function
	done := make(chan bool)

	// write messages to client goroutine
	go client.writeMessages(ctx, done)

	<-done

}

// takes in a send only done channel as paramater to signal parent goroutine
func (c *Client) writeMessages(ctx context.Context, done chan<- bool) {

	defer func() {
		c.Conn.Close()
		log.Println("CANCELLING CONTEXT HERE")
		done <- true
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case message, ok := <-c.Send:

			if !ok {
				log.Printf("Send channel close for client %s", c.ClientID.String())
				return
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, message)

			if err != nil {
				log.Printf("Websocket write error: %v", err)
				return
			}
			log.Printf("Message sent to client %s: %s", c.ClientID.String(), string(message))

		}
	}

}

func Broadcast(ctx context.Context) {

	for message := range manager.Send {
		for _, client := range manager.Clients {

			// send message to client's send channel
			client.Send <- message

		}
	}

}

func GetRealTimePrices(ctx context.Context) {

	// gets realtime time price data every 5 mins and writes to the managers Send channel

	for {
		time.Sleep(10 * time.Second)

		for _, coin_name := range api.CoinNames {
			manager.mu.Lock()
			// api call for real-time price data
			log.Printf("Api call for realtime price of %v", coin_name)
			data, err := json.Marshal(api.RealTimePriceAPI(coin_name))

			if err != nil {
				log.Printf("Error getting realtime price for %v", coin_name)
				data = []byte("")
			}


			manager.Send <- data
			manager.mu.Unlock()

		}

	}

}
