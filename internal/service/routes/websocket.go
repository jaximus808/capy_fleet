package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jaximus808/capy_websocket/internal/service/multiplayer"
)

// finish vid here
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//in the future make this the actual webserver domain
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		return origin == "http://127.0.0.1:8080" || origin == "http://localhost:8080"
	},
}

func create_websocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error making websocket")
		return
	}
	defer conn.Close()

	fmt.Println("connected?")
	conn_id := multiplayer.AddClient(conn)
	defer multiplayer.RemoveClient(conn_id)

	//lol wtf is this

	for {
		//messages should not include the author id, as they will be known by the server
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("client disconnected of id: ", conn_id, err)
			return
		}
		if mt == websocket.BinaryMessage {
			multiplayer.HandleMessage(message, conn_id)
		}
	}
}

// I need to mkae a multiplayer server to game oinstance bridge that allows me to talk to each other without the need to byte convert
