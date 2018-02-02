package p2p

import (
	"common"
)

type ServiceType uint64

const (
	UNNAMED ServiceType = 0x00

	NODE_NETWORK ServiceType = 0x01
)

//The version message provides information about the transmitting node to the receiving node at the beginning of a connection.
//reference : https://bitcoin.org/en/developer-reference#Version
type Version_Message struct {
	Header message_header
	//The highest protocol version understood by the transmitting node.
	Version int32

	//The services supported by the transmitting node encoded as a bitfield.
	Services uint64

	//The current Unix epoch time according to the transmitting node’s clock.
	TimeStamp int64

	RecipientAddress NetworkAddress

	SenderAddress NetworkAddress

	//A random nonce which can help a node detect a connection to itself. If the nonce is 0, the nonce field is ignored. If the nonce is anything else, a node should terminate the connection on receipt of a version message with a nonce it previously sent.
	Nonce uint64

	//required if user agent byte > 0,
	UserAgent string

	StartHeight int32

	//Transaction relay flag. If 0x00, no inv messages or tx messages announcing new transactions should be sent to this client until it sends a filterload message or filterclear message. If the relay field is not present or is set to 0x01, this node wants inv messages and tx messages announcing new transactions.
	Relay bool
}

func (versionMsg *Version_Message) Decode(payload []byte) {
	versionMsg.Header.init(VERSION, payload)
	input := common.NewBitcoinInput(payload)
	input.ReadNum(&versionMsg.Version)
	input.ReadNum(&versionMsg.Services)
	input.ReadNum(&versionMsg.TimeStamp)
	versionMsg.RecipientAddress.Init(input, true)
	versionMsg.SenderAddress.Init(input, true)
	input.ReadNum(&versionMsg.Nonce)

	userAgentBytesCount, _ := input.ReadVarInt()
	userAgentBytes := make([]byte, userAgentBytesCount)
	input.ReadBytes(userAgentBytes)
	versionMsg.UserAgent = string(userAgentBytes)

	input.ReadNum(&versionMsg.StartHeight)
	input.ReadNum(&versionMsg.Relay)
}
