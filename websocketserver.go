package WebsocketServer

import "fmt"

type WebsocketServer struct {
	//Register channel will broadcast to all connected users when a new user has joined.
	Register chan *client

	//Unregister channel will broadcast to all connected users when a user disconnects from the server.
	Unregister chan *client

	//Broadcast will accept messages from a client, and send it each connected client.
	Broadcast chan message

	//Clients will store a reference to each connected client on the socket server.
	Clients map[*client]bool

    //Rooms will map the name of the room as a string to the clients joined in that room.
    Rooms map[string]*client
}

//Returns a new instance of a WebsocketServer with all of its values initialized.
func NewSocketServer () *WebsocketServer{
	return &WebsocketServer{
		Register:   make(chan *client),
		Unregister: make(chan *client),
		Broadcast:  make(chan message),
		Clients:    make(map[*client]bool),
		Rooms:      make(map[string]*client),
	}
}

//Function to initialize the server, and allow it to process the clients. It must be run in a seperate goroutine.
func (ws *WebsocketServer) Start() {
	for {
		select {
			case client := <-ws.Register:
				ws.Clients[client] = true
				fmt.Println("num users:", len(ws.Clients))
				ws.broadcastMessage(message{Body: fmt.Sprintf("Client %s connected...", client.ID), Type: 2, ClientID: client.ID})
				break
			case client := <-ws.Unregister:
				delete(ws.Clients, client)
				fmt.Println("num users:", len(ws.Clients))
				fmt.Println("client", client, "disconnected")
				ws.broadcastMessage(message{Body: fmt.Sprintf("Client %s disconnected...", client.ID), Type: 2, ClientID: client.ID})
				break
			case message := <-ws.Broadcast:
				ws.broadcastMessage(message)
		}
	}
}

func (ws *WebsocketServer) broadcastMessage(messageToSend message){
	for client := range ws.Clients {
		if client.ID != messageToSend.ClientID{
			client.Conn.WriteJSON(messageToSend)
		}
	}
}