package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// finish vid here
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func create_websocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error making websocket")
		return
	}
	defer conn.Close()
	conn.WriteMessage(websocket.TextMessage, []byte("Hello!!!!"))
}
