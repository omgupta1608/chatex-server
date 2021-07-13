package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		event:      make(chan []byte),
	}
}

var hub *Hub

func init() {
	hub = createHub()
	// init the Hub
	go hub.run()
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn, hub: hub, send: make(chan []byte, 256)}

	// readPump and writePump
	go client.handleRead()
	go client.handleWrite()
}
