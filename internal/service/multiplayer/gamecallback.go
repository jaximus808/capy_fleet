package multiplayer

import (
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/jaximus808/capy_websocket/internal/service/bridge"
)

func HandleEvent() {
	event_bus.Subscribe("packet_send", SendPacket)
}

func SendPacket(event *bridge.Event) error {
	websocket_mut.Lock()
	defer websocket_mut.Unlock()
	data := *event.GetData()
	action, action_exist := data["packet"].(bridge.Action)

	if !action_exist {
		return errors.New("invalid data")
	}

	targets := action.GetTargets()

	//need to put flags for debug mode
	//make an enum for this
	//TODO: fix jitterring
	switch action.GetSpecialTarget() {

	case 0:
		for _, target := range targets {
			client, client_exist := clients[target]

			if !client_exist {
				continue
			}

			fmt.Println("sending packet to: ", target)
			client.conn.WriteMessage(websocket.BinaryMessage, action.GetPacket().GetBuffer())
		}
	case 1:
		for _, client := range clients {
			client.conn.WriteMessage(websocket.BinaryMessage, action.GetPacket().GetBuffer())
		}
	case 2:

		ignore_set := make(map[uint]struct{})

		for _, ignored_ids := range action.GetTargets() {
			ignore_set[ignored_ids] = struct{}{}
		}

		for cid, client := range clients {
			_, has_client := ignore_set[cid]
			if !has_client {
				client.conn.WriteMessage(websocket.BinaryMessage, action.GetPacket().GetBuffer())
			}
		}

	}

	//check if the player has already joined

	return nil
}
