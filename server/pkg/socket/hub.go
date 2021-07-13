package socket

import (
	"encoding/json"
	"errors"

	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

type (
	Hub struct {
		clients map[*Client]bool

		register chan *Client

		unregister chan *Client

		event chan []byte
	}
)

var (
	Events = map[string]func(*Hub, map[string]interface{}) error{
		types.NEW_MESSAGE:    NewMessageHandler,
		types.DELETE_MESSAGE: DeleteMessageHandler,
		types.TYPING:         TypingHandler,
		types.USER_CONN:      UserConnHandler,
	}
)

func createHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		event:      make(chan []byte),
	}
}

/*
 * Handle the incoming event
 * Identify the event type and its respective handler
 * Execute the handler
 */
func (hub *Hub) handleEvent(data []byte) {
	event, payload, err := getEventAndData(data)
	if err != nil {
		// log error
		exception.LogError(err, "cannot get event", exception.INTERNAL_SOCKET_ERROR)
		// close connection

	}
	// execute respective handler
	err = Events[event](hub, payload)

	if err != nil {
		// log error
		exception.LogError(err, err.Error(), exception.INTERNAL_SOCKET_ERROR)
		// close connection

	}
}

func getEventAndData(data []byte) (string, map[string]interface{}, error) {
	var msg interface{}
	if err := json.Unmarshal(data, &msg); err != nil {
		// log error
		exception.LogError(err, "cannot unmarshal payload", exception.INTERNAL_SOCKET_ERROR)
		return "", nil, err
	}
	msgMap := msg.(map[string]interface{})
	if msgMap["name"] == "" {
		// log error (no event)
		exception.LogError(errors.New("no event provided"), "missing field \"event\"", exception.INTERNAL_SOCKET_ERROR)
		return "", nil, errors.New("no event provided")
	}

	return msgMap["name"].(string), msgMap, nil
}

func (hub *Hub) run() {
	for {
		select {
		// handle client registration
		case client := <-hub.register:
			hub.clients[client] = true

		// handle client unregistration (disconnection)
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}

		// handle incoming event
		case data := <-hub.event:
			hub.handleEvent(data)
		}
	}
}
