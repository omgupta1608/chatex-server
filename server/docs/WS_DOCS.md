# Web Socket Docs

The implementation is quite simple and easy to understand just by looking at the code. 
- Prerequisites 
  - Go Routines 
  - Channels

## Important points
- Connection to the socket server can be made by simply connecting to the ` /ws ` path using the standard Web Socket Protocol.
- Connections are unauthenticated as of now.
- As ` gorilla/websocket ` doesn't provide any inbuilt mechanism to emit and recieve events on the basis of event names, we have to implement it ourselves.(both on the client and server side). Refer [#24](https://github.com/omgupta1608/chatex/issues/24)