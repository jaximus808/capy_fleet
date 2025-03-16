package game

import "github.com/jaximus808/capy_websocket/internal/service/bridge"

// packet format:
// packect_id: int32 client_join : 0
// client id: int64
// client name: string
// startingpos_x: Float64
// startingpos_y: Float64
func createJoinPacket(player *Player) *bridge.Packet {

	new_packet := bridge.CreatePacket(1024)
	new_packet.WriteInt32(int32(client_join))
	new_packet.WriteInt64(int64(player.client_id))
	new_packet.WriteString(player.uname)
	//make a write vec 2 functgion
	new_packet.WriteFloat64(player.pos.X())
	new_packet.WriteFloat64(player.pos.Y())
	return new_packet
}

// packet format:
// packet id: int32 leaving : 1
// client id: int64
func createLeavePacket(id uint) *bridge.Packet {

	new_packet := bridge.CreatePacket(1024)
	new_packet.WriteInt32(int32(client_disconnect))
	new_packet.WriteInt64(int64(id))
	return new_packet
}

// packet format:
// packet id: int32 leaving : 1
// client id: int64
// client x: float64
// client y: float64
func createMovePacket(player *Player) *bridge.Packet {

	new_packet := bridge.CreatePacket(1024)
	new_packet.WriteInt32(int32(player_move))
	new_packet.WriteInt64(int64(player.client_id))
	new_packet.WriteFloat64(player.pos.X())
	new_packet.WriteFloat64(player.pos.Y())
	return new_packet
}
