package multiplayer

import (
	"github.com/gorilla/websocket"
	"github.com/jaximus808/capy_websocket/internal/service/game"
)

type Client struct {
	conn *websocket.Conn
	id   uint
	//user_id when we get webo auth going
	player *game.Player
}

func CreateClient(conn *websocket.Conn, id uint) *Client {
	return &Client{
		conn: conn,
		id:   id,
	}
}

func (client *Client) SetPlayer(player *game.Player) {
	client.player = player
}
