package game

import (
	"errors"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
	"github.com/jaximus808/capy_websocket/utils"
)

const (
	ping uint = iota
	move
)

var gameplay_inputs = map[uint]func(*bridge.Packet, uint) error{

	ping: Pong,
}

func Pong(packet *bridge.Packet, author_id uint) error {
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

	player.target_pos.Set(target_vec)
	return nil
}
