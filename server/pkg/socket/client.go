package socket

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"github.com/omgupta1608/chatex/server/pkg/exception"
)

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

const (
	// time allowed to write a message
	writeWait = 10 * time.Second

	// time allowed to read the next pong message
	pongWait = 60 * time.Second

	// send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// maximum message size allowed
	maxMessageSize = 1024 // configurable
)

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

func newClient(conn *websocket.Conn) *Client {
	return &Client{conn: conn, hub: hub, send: make(chan []byte, 256)}
}

/*
 * Code for handleRead and handleWrite is used from the standard readPump and writePump implementation.
 * No business logic will be included these function.
 * All logic will be performed on the Hub.
 * Client is only used to pump messages to and from the hub
 */

// pumps messages from the web socket connection to the hub
func (c *Client) handleRead() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				exception.LogError(err, "Error in reading message from socket connection", exception.INTERNAL_SOCKET_ERROR)
			}
			break
		}
		message = bytes.TrimSpace(message)
		c.hub.event <- message
	}
}

// pumps message from the hub to the web socket connection
func (c *Client) handleWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write(<-c.send)
			// }

			if err := w.Close(); err != nil {
				exception.LogError(err, "Error in writing message from Hub to Websocket Conn", exception.INTERNAL_SOCKET_ERROR)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				exception.LogError(err, "Error in writing message from Hub to Websocket Conn", exception.INTERNAL_SOCKET_ERROR)
				return
			}
		}
	}
}
