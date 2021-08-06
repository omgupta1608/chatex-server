package socket

import (
	"encoding/json"
	"strconv"

	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/types"
	"github.com/omgupta1608/chatex/server/pkg/utils"
)

/*
 * The code in the first three functions are redundant but I think its better to keep them separate,
   just in case something needs to be implemented for just one of them in future.
*/

func NewMessageHandler(hub *Hub, payload *types.EventFormat) error {
	scid := payload.Data.S_ID + "." + payload.Data.R_ID + "." + strconv.FormatInt(utils.GetCurrentUnixTimeStamp(), 10)
	ts := utils.GetCurrentUnixTimeStamp()
	for client := range hub.clients {
		// send reciever the new message
		if client.Id == payload.Data.R_ID {
			// set message id
			payload.Data.SERVER_CHAT_ID = scid
			// set timestamp
			payload.Data.ServerTimestamp = ts
			rByte, err := json.Marshal(payload)
			if err != nil {
				// log error
				exception.LogError(err, "unable to marshal data", exception.INTERNAL_SOCKET_ERROR)
				return err
			}
			client.send <- rByte
		}
		// pong sender with (scid and ccid)
		if client.Id == payload.Data.S_ID {
			response := &types.EventFormat{
				Event_Name: "NEW_MSG_PONG",
				Data: types.Message{
					CLIENT_CHAT_ID:  payload.Data.CLIENT_CHAT_ID,
					SERVER_CHAT_ID:  scid,
					S_ID:            "SERVER",
					R_ID:            payload.Data.S_ID,
					Content:         "",
					ServerTimestamp: ts,
					Message_Type:    "event",
					Data:            nil,
				},
			}
			rByte, err := json.Marshal(response)
			if err != nil {
				// log error
				exception.LogError(err, "cannot marshal socket response", exception.INTERNAL_SOCKET_ERROR)
				return err
			}
			client.send <- rByte
		}
	}
	return nil
}

func DeleteMessageHandler(hub *Hub, payload *types.EventFormat) error {
	for client := range hub.clients {
		if client.Id == payload.Data.R_ID {
			// set timestamp
			payload.Data.ServerTimestamp = utils.GetCurrentUnixTimeStamp()
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
			// set timestamp
			payload.Data.ServerTimestamp = utils.GetCurrentUnixTimeStamp()
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

// This function is invoked whenever a client connects to the server. All the messages that were designated to the client but not recieved by it will get delivered as an array of messages
func GetUnDeliveredMessages(hub *Hub, client *Client) ([]types.Message, error) {
	// get UD_MSGS from redis
	return nil, nil
}
