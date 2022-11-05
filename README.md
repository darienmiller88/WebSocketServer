## Websocket Server implemenation using Gorilla Websockets

<hr />

#### This is a slightly modified websocket server implementation of the base example given by the team that created Gorilla websockets, with some personal touches added.

<hr />

### Features
- Server will be able to broadcast messages to every connected client, or every client except the original sender.
- Rooms will be implemented soon, allowing users to send messages only to other users in the same room

### Installation
```
go get https://github.com/darienmiller88/WebSocketServer
```