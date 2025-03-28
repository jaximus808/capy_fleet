package multiplayer

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jaximus808/capy_websocket/internal/service/bridge"
	"github.com/jaximus808/capy_websocket/internal/service/game"
)

var clients map[uint]*Client = map[uint]*Client{}
var out_queue bridge.Queue = *bridge.CreateQueue(256) //multi server writes to this
//var in_queue bridge.Queue = *bridge.CreateQueue(256) //deprecated prob dont need // multi server reads this

var event_bus bridge.EventBus = *bridge.CreateEventBus() //this handles data sent between threads that are server auth actions, such as connecting etc
var websocket_mut = sync.Mutex{}

// server id is reserved as 0
var cur_id uint = 1

func SpinUpMultiPlayerGame() {
	// lol scary asf
	HandleEvent()
	go game.SpinUpGame(&out_queue, &event_bus)
}

func AddClient(conn *websocket.Conn) uint {
	clients[cur_id] = CreateClient(conn, cur_id)

	//tell game server to add the player

	welcome_msg := bridge.CreateEvent()
	welcome_msg.Add("uid", cur_id)

	err := event_bus.Publish("welcome_msg", welcome_msg)

	if err != nil {
		fmt.Println(err.Error())
	}
	//event_bus.Publish("create_user", create_event)

	cur_id++
	return cur_id - 1
}
func RemoveClient(cur_id uint) {
	create_event := bridge.CreateEvent()
	create_event.Add("uid", cur_id)
	event_bus.Publish("disconnect_user", create_event)
	delete(clients, cur_id)
}

func HandleMessage(msg []byte, client_id uint) {
	//again i need to mkae a packet size constant
	packet := bridge.ConvertToPacket(msg, 1024)

	action := bridge.CreateAction(client_id, packet)
	out_queue.AddAction(action)
}
