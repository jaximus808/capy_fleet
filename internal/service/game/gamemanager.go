package game

import (
	"fmt"
	"time"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
)

var players map[uint]*Player

const SERVER_ID uint = 0

var eventbus *bridge.EventBus
var tick_rate = time.Second / 30

func SpinUpGame(action_queue *bridge.Queue, event_bus *bridge.EventBus) {

	fmt.Println("spinning up game server")
	eventbus = event_bus
	HandleEvents()
	//tick rate set to 30, maybe make this a parameter
	players = make(map[uint]*Player)
	ticker := time.NewTicker(tick_rate)
	defer ticker.Stop()

	for range ticker.C {
		if action_queue.Peak() {
			action, err := action_queue.ProcessAction()
			if err != nil {
				fmt.Printf("soemthing went SUPER wrong")
				return
			}
			ActionCallback(&action)
		}
		for _, player := range players {
			player.Update(tick_rate)
		}
	}
}

func BroadcastToClient(client_id uint, packet *bridge.Packet) {
	action := bridge.CreateAction(SERVER_ID, packet)
	action.AddTarget(client_id)
	err := eventbus.Publish("packet_send", bridge.CreateEventAction(*action))
	if err != nil {
		fmt.Println(err.Error())
	}

}
func BroadcastToClients(client_ids []uint, packet *bridge.Packet) {
	action := bridge.CreateAction(SERVER_ID, packet)
	action.AddTargets(client_ids)
	eventbus.Publish("packet_send", bridge.CreateEventAction(*action))
}

func BroadcastToAll(packet *bridge.Packet) {
	action := bridge.CreateAction(SERVER_ID, packet)

	action.SetSpecial(1)
	eventbus.Publish("packet_send", bridge.CreateEventAction(*action))
}

// cant i just combine these functions?
func BroadcastToAllExcept(ignored_id uint, packet *bridge.Packet) {
	action := bridge.CreateAction(SERVER_ID, packet)

	action.AddTarget(ignored_id)
	action.SetSpecial(2)
	eventbus.Publish("packet_send", bridge.CreateEventAction(*action))
}
func BroadcastToAllExceptMulti(ignored_ids []uint, packet *bridge.Packet) {
	action := bridge.CreateAction(SERVER_ID, packet)

	action.AddTargets(ignored_ids)
	action.SetSpecial(2)
	eventbus.Publish("packet_send", bridge.CreateEventAction(*action))
}
