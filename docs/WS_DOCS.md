# Web Socket Docs

The implementation is quite simple and easy to understand just by looking at the code. 
- Prerequisites 
  - Go Routines 
  - Channels

## Important points
- Connection to the socket server can be made by simply connecting to the ` /ws ` path using the standard Web Socket Protocol.
- Connections are unauthenticated as of now.
- As ` gorilla/websocket ` doesn't provide any inbuilt mechanism to emit and recieve events on the basis of event names, we have to implement it ourselves.(both on the client and server side). Refer [#24](https://github.com/omgupta1608/chatex/issues/24)

## Event/Payload (Payload struct is same for all requests, only the data inside it changes)
- Events
  - ` NEW_MESSAGE `
    - When sending a new message 
  - ` DELETE_MESSAGE `
    - When deleting a message
  - ` TYPING `
    - When typing
  - ` USER_CONN `
    - User connect/disconnect event for updating last seen and online status
- Payload Structure
  ```
  "name":         Name of the event (NEW_MESSAGE | DELETE_MESSAGE | TYPING | USER_CONN),
  "data": {
      "scid":        Server Chat Id (Id of new message | Id of message to delete | "" in case of typing and user_conn),
      "ccid":       Client Chat Id (Id of new message | Id of message to delete | "" in case of typing and user_conn),
      "sid":        Sender Id ("SERVER" in case on payload being sent by the server),
      "rid":        Reciever Id,
      "content":    Content of the message ("" is case of typing, user_conn and delete msg),
      "sts" :        Server Unix Timestamp (should be left empty, as it will be attached on the server to avoid conflicts), [Format -> MM-DD-YYYY::hh:mm:ss]
      "m_type":     Type of the message ("text", "img", "video", "file", "doc" etc. | "event" is case of typing, user_conn and delete msg)
  }
  ```
