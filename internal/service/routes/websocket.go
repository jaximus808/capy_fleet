package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	it := 0
	for {
		it++
		conn.WriteMessage(websocket.TextMessage, []byte("Hello!!!!"))
		time.Sleep(time.Second)
		if it == 300 {
			break
		}
	}
}
