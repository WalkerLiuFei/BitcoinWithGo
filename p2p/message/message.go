package message

import (
	"constants"
	"github.com/pkg/errors"
	"strings"
)

type Message interface {
	Decode([]byte)

	Encode() []byte

	GetPayload() []byte
}

func DecodeMessage(inputStream []byte) (Message, error) {
	if len(inputStream) == 0 {
		return nil, errors.New("empty message")
	}
	header := decodeHeader(inputStream)
	var message Message

	switch header.magic {
	case constants.MAIN_NET.Magic,
		constants.TEST_NET.Magic,
		constants.REG_TEST_NET.Magic:
	default:
		return nil, errors.New("bad magic byte")
	}
	msgType := constants.MessageType(strings.TrimRightFunc(string(header.command), func(r rune) bool {
		return !(r > 96 && r < 123)
	}))
	switch msgType {
	case constants.HEADER:
		message = &HeaderMessage{}
	case constants.REJECT_MESSAGE:
		message = &Reject_Message{}
	case constants.MERKLE_BLOCK:
		message = &MerkleBlockMessage{}
	case constants.INV_MESSAGE:
		message = &InvMessage{}
	case constants.PONG_MESSAGE:
		message = &PongMessage{}
	case constants.PING_MESSAGE:
		message = &PingMessage{}
	case constants.NOT_FOUND, constants.VAERSION_ACK, constants.GET_ADDRESS:
		message = &NonePayloadMessage{}
	case constants.ADDRESS_MESSAGE:
		message = &Address_Message{}
	case constants.MEM_POOL_MESSAGE:
	case constants.VERSION:
		message = &version_Message{}
	}
	//前24个字节为header ， header后面的为payload
	message.Decode(inputStream)
	return message, nil
}
