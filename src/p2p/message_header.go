package p2p

import (
	"go.uber.org/zap"
	"utils"
)

type Message interface {
	GetPayload() []byte

	Init([]byte)

	GetBytes() []byte
}

// reference : https://bitcoin.org/en/developer-reference#message-headers
type message_header struct {
	//
	MAGIC []byte
	//标识信息的类型,size 固定为12
	Command []byte
	//Payload size
	PayloadSize uint32
	//First 4 bytes of SHA256(SHA256(payload)) in
	CheckSum []byte
}

var logger, _ = zap.NewProduction()

func (msg *message_header) init(cmd string, payload []byte) {
	msg.MAGIC = []byte{0Xf9, 0Xbe, 0Xb4, 0Xd9}
	msg.Command = []byte(cmd)
	if len(msg.Command) < 12 {
		index := len(msg.Command)
		for index < 12 {
			msg.Command = append(msg.Command, 0)
		}
	}
	msg.PayloadSize = uint32(len(payload))
	msg.CheckSum = utils.DoubleHash(payload)[:4]
}
