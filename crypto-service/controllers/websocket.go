package controllers

import (
	"sync"

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
