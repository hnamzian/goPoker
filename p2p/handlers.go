package p2p

import (
	"fmt"
	"io"
)

type Handler interface {
	HandleMessage(*Message) error
}

type Defaulthandler struct {}

func (h *Defaulthandler) HandleMessage(msg *Message) error {
	buff, err := io.ReadAll(msg.payload)
	if err != nil {
		return fmt.Errorf("read message payload failed: %s", err)
	}
	
	fmt.Printf("Message from (%s): %s\n", msg.from, string(buff))

	return nil
}
