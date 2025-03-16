package multiplayer

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/jaximus808/capy_websocket/internal/service/bridge"
	"github.com/jaximus808/capy_websocket/internal/service/game"
)

var clients map[uint]*Client = map[uint]*Client{}
var out_queue bridge.Queue = *bridge.CreateQueue(256) //multi server writes to this
//var in_queue bridge.Queue = *bridge.CreateQueue(256) //deprecated prob dont need // multi server reads this

var event_bus bridge.EventBus = *bridge.CreateEventBus() //this handles data sent between threads that are server auth actions, such as connecting etc

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
	create_event := bridge.CreateEvent()
	create_event.Add("uid", cur_id)
	create_event.Add("name", "test_"+strconv.Itoa(int(cur_id)))

	test_welcome := bridge.CreateEvent()
	test_welcome.Add("uid", cur_id)

	err := event_bus.Publish("welcome_msg", test_welcome)

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
