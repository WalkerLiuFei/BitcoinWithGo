package p2p

import "go.uber.org/zap"
type Message struct {
	Command []byte
}
var logger,_ = zap.NewProduction()
func (msg *Message) Init(cmd string) {
	msg.Command = []byte(cmd)
}

func (msg *Message) getCommandFrom(cmd []byte)  {
	length := len(cmd)
	for length >= 0{

	}
}