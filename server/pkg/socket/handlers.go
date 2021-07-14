package socket

import (
	"encoding/json"

	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

/*
 * The code in the first three functions are redundant but I think its better to keep them separate,
   just in case something needs to be implemented for just one of them in future.
*/

func NewMessageHandler(hub *Hub, payload *types.EventFormat) error {
	for client := range hub.clients {
		if client.Id == payload.Data.R_ID {
			data, err := json.Marshal(payload)
			if err != nil {
				// log error
				exception.LogError(err, "unable to marshal data", exception.INTERNAL_SOCKET_ERROR)
				return err
			}
			client.send <- data
		}
	}
	return nil
}

func DeleteMessageHandler(hub *Hub, payload *types.EventFormat) error {
	for client := range hub.clients {
		if client.Id == payload.Data.R_ID {
			data, err := json.Marshal(payload)
			if err != nil {
				// log error
				exception.LogError(err, "unable to marshal data", exception.INTERNAL_SOCKET_ERROR)
				return err
			}
			client.send <- data
		}
	}
	return nil
}

func TypingHandler(hub *Hub, payload *types.EventFormat) error {
	for client := range hub.clients {
		if client.Id == payload.Data.R_ID {
			data, err := json.Marshal(payload)
			if err != nil {
				// log error
				exception.LogError(err, "unable to marshal data", exception.INTERNAL_SOCKET_ERROR)
				return err
			}
			client.send <- data
		}
	}
	return nil
}

/*
 * Whenever the user disconnects, the last seen is updated to the current timestamp
 * Whenever the user connects, first thing to be done is to delete its last seen from the cache (no last seen means user is online)
 */
func UserConnHandler(hub *Hub, payload *types.EventFormat) error {
	// update redis for last seen
	return nil
}
