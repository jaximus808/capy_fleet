package game

import (
	"errors"
	"fmt"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
	"github.com/jaximus808/capy_websocket/utils"
)

const (
	ping uint = iota
	ready
	move
)

var gameplay_inputs = map[uint]func(*bridge.Packet, uint) error{

	ping:  Pong,
	ready: Ready,
	move:  NewMoveTarget,
}

func Pong(packet *bridge.Packet, author_id uint) error {
	return nil
}
func Ready(packet *bridge.Packet, author_id uint) error {
	uname, err := packet.ReadString()
	if err != nil {
		return errors.New("cant read name")
	}
	_, id_exhausted := players[author_id]
	if id_exhausted {
		return errors.New("player already added")
	}
	fmt.Println("Creating user with name:", uname)
	new_player := CreatePlayer(uname, author_id, "starting", 200.0, 200.0)

	players[author_id] = new_player
	//need to make broadcast

	new_join_packet := createJoinPacket(new_player)

	new_create_packet := createGameInfoPacket(author_id)
	BroadcastToClient(author_id, new_create_packet)

	BroadcastToAllExcept(author_id, new_join_packet)

	return nil
}

// make a player speed, and then move the player towrads the target at a set speed lol
func NewMoveTarget(packet *bridge.Packet, author_id uint) error {

	target_pos_x, err1 := packet.ReadFloat64()
	target_pos_y, err2 := packet.ReadFloat64()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	player, has_player := players[author_id]

	if !has_player {
		return errors.New("player to move does not mf exist???")
	}

	target_vec := utils.Vector2(target_pos_x, target_pos_y)

	player.SetNewMoveVector(target_vec)
	return nil
}
