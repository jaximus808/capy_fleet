package game

import (
	"errors"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
)

// might need a state for "ready lol"

func HandleEvents() {

	eventbus.Subscribe("welcome_msg", WelcomeMsg)
	eventbus.Subscribe("disconnect_user", RemoveUser)

}

func WelcomeMsg(event *bridge.Event) error {

	data := *event.GetData()
	id, id_exists := data["uid"].(uint)

	if !id_exists {
		return errors.New("user doesn't exists")
	}
	welcome_packet := createWelcomePacket(int64(id))

	BroadcastToClient(id, welcome_packet)
	return nil

}

func RemoveUser(event *bridge.Event) error {
	data := *event.GetData()
	id, id_exists := data["uid"].(uint)

	if !id_exists {
		return errors.New("invalid data")
	}
	_, player_exists := players[id]
	if !player_exists {
		return errors.New("player does not exist")
	}
	//create a disconect handler like if their in an event etc
	delete(players, id)
	//need to make broadcast

	new_packet := createLeavePacket(id)

	BroadcastToAll(new_packet)
	return nil
}
