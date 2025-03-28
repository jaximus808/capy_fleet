package game

import (
	"fmt"
	"time"

	"github.com/jaximus808/capy_websocket/utils"
)

type Player struct {
	uname      string
	client_id  uint
	location   string
	pos        *utils.Vec2
	target_pos *utils.Vec2
	speed      float64
	moving     bool
}

func CreatePlayer(uname string, client_id uint, location string, x_pos float64, y_pos float64) *Player {
	return &Player{
		uname:      uname,
		client_id:  client_id,
		location:   location,
		pos:        utils.Vector2(x_pos, y_pos),
		target_pos: utils.Vector2(x_pos, y_pos),
		speed:      100, //5 mps
		moving:     false,
	}
}

func (p *Player) SetNewMoveVector(new_pos *utils.Vec2) {
	if p.pos.Equals(new_pos) {
		return
	}
	p.target_pos = new_pos
	p.moving = true
}

func (p *Player) Update(deltaTime time.Duration) {

	if p.moving {

		initial_dir := p.pos.VecTowards(p.target_pos)
		// fmt.Println("MOvING")
		// fmt.Println(float64(deltaTime.Seconds()))
		// fmt.Println(initial_dir.X())
		// fmt.Println(initial_dir.Y())
		vel_vec := initial_dir.Scalec(p.speed * float64(deltaTime.Seconds()))

		// fmt.Println(vel_vec.X())
		// fmt.Println(vel_vec.Y())
		p.pos.Add(vel_vec)

		//if the direction has changed, meaning the movement made the capy go past the point, then we've passed it
		if p.pos.Dist(p.target_pos) < 0.5 || vel_vec.Mag() == 0 {

			fmt.Println("made it!")
			fmt.Println(p.target_pos)
			p.pos.Set(p.target_pos)
			p.moving = false
			fmt.Println(p.target_pos)
		}
		movement_packet := createMovePacket(p)
		BroadcastToAll(movement_packet)
	}
}
