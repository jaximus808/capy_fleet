package game

import (
	"errors"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
)

// might need a state for "ready lol"
const (
	client_join uint = iota
	client_disconnect
	player_move
)

func HandleEvents() {

	eventbus.Subscribe("welcome_msg", WelcomeMsg)
	eventbus.Subscribe("create_user", CreateUser)
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

func CreateUser(event *bridge.Event) error {

	data := *event.GetData()
	id, id_exists := data["uid"].(uint)

	uname, uname_exist := data["uname"].(string)

	if !id_exists || !uname_exist {
		return errors.New("invalid data")
	}
	//check if the player has already joined
	_, id_exhausted := players[id]
	if id_exhausted {
		return errors.New("player already added")
	}

	new_player := CreatePlayer(uname, id, "starting", 200.0, 200.0)
	players[id] = *new_player
	//need to make broadcast
	new_packet := createJoinPacket(new_player)

	BroadcastToAll(new_packet)

	//should I send all the game info here?
	// I think maybe make a seperate event to give game info to the player
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
