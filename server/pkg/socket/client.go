package socket

import "github.com/gorilla/websocket"

type (
	Client struct {
		// the main hub instance
		hub *Hub

		// the actual socket connection
		conn *websocket.Conn

		// []byte channel for sending and recieving data to and from the hub
		send chan []byte
	}
)

func (c *Client) handleRead() {
}

func (c *Client) handleWrite() {

}
