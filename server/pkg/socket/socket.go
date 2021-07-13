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

	// create new client
	client := newClient(conn)

	// register client
	client.hub.register <- client

	// readPump and writePump
	go client.handleRead()
	go client.handleWrite()
}
