package server

import (
	"log"
	"sync"
)

type Hub struct {
	Join      chan *Client
	Broadcast chan Message
	Leave     chan *Client
	Clients   map[*Client]bool
	Mutex     sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Join:      make(chan *Client),
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan Message),
		Leave:     make(chan *Client),
	}

}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Join:
			log.Printf("Client %v joined!", client.Username)

			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()

		case msg := <-h.Broadcast:
			log.Println("Message Broadcast: ", msg)
			h.Mutex.Lock()
			for client := range h.Clients {
				client.WriteMessage(msg)
			}
			h.Mutex.Unlock()
		}
	}
}
