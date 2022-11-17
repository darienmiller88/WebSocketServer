package WebsocketServer

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

//Function to serve the end point "/ws" and process the clients messages.
func ServeWebSocketServer(ws *WebsocketServer, res http.ResponseWriter, req *http.Request) {
	//Upgrade the server to be work with websockets.
	conn, err := upgrade(res, req)

	if err != nil {
		panic(err)
	}

	client := &client{
		ID:   uuid.NewString(),
		Conn: conn,
	}

	//Send the newly created client to the register channel so they can be sent to each client
	ws.Register <- client
	fmt.Println("client connected:", client)
	client.processMessage(ws)
}