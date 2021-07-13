package socket

type (
	Hub struct {
		clients map[*Client]bool

		register chan *Client

		unregister chan *Client

		event chan []byte
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
func (h *Hub) handleEvent(event []byte) {

}

func (h *Hub) run() {
	for {
		select {
		// handle client registration
		case client := <-h.register:
			h.clients[client] = true

		// handle client unregistration (disconnection)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		// handle incoming event
		case event := <-h.event:
			h.handleEvent(event)
		}
	}
}
