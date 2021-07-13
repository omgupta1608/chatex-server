package socket

type (
	Hub struct {
		clients map[*Client]bool

		register chan *Client

		unregister chan *Client

		event chan []byte
	}
)

func (h *Hub) run() {

}
