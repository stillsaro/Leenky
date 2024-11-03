package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Hub      *Hub
	Username string
	WsConn   *websocket.Conn
	Send     chan Message
}

type Message struct {
	Sender  string `json:"sender"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

func (s *Server) LeenkyChatWsHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := Client{
		Hub:      s.Hub,
		Username: c.Get("username").(string),
		WsConn:   conn,
		Send:     make(chan Message),
	}
	s.Hub.Join <- &client

	client.ReadMessage()
	return nil
}

func (c *Client) ReadMessage() {
	defer c.WsConn.Close()
	for {
		_, content, err := c.WsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := Message{
			Sender:  c.Username,
			Time:    time.Now().Format("15:04"),
			Content: string(content),
		}
		c.Hub.Broadcast <- msg
	}
}

func (c *Client) WriteMessage(msg Message) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}
	if err = c.WsConn.WriteMessage(websocket.TextMessage, jsonMsg); err != nil {
		log.Printf("Error: Faild to write the message to websocket connection: %v", err)
	}

	log.Printf("New message from \"%v\"at '%v': '%v'", msg.Sender, msg.Time, msg.Content)
}
