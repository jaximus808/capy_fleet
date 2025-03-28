package game

import (
	"errors"
	"fmt"

	"github.com/jaximus808/capy_websocket/internal/service/bridge"
)

// actual gameplay input to output
func ActionCallback(action *bridge.Action) error {

	author_id := action.GetAuthor()

	packet := action.GetPacket()

	gameinput_id, err_actid := packet.ReadInt32()

	if err_actid != nil {
		return err_actid
	}

	gameplay_func, valid_actid := gameplay_inputs[uint(gameinput_id)]

	if !valid_actid {
		return errors.New("invalid packet id")
	}

	gameplay_err := gameplay_func(packet, author_id)
	if gameplay_err != nil {
		fmt.Println(gameplay_err.Error())
	}
	return gameplay_err

}
