package WebsocketServer

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgrade(res http.ResponseWriter, req *http.Request) (*websocket.Conn, error) {
	websocketConnection, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Println(err)
		return websocketConnection, err
	}

	return websocketConnection, nil
}