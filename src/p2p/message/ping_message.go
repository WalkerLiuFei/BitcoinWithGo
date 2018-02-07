package message

import (
	"common"
	"math/rand"
	"constants"
)

//https://bitcoin.org/en/developer-reference#ping
type PingMessage struct {
	Header message_header
	//random number to confirm connection
	Nonce uint64
}

func (ping *PingMessage) Init([]byte) {
	ping.Nonce = uint64(rand.Int63())
	ping.Header.init(constants.PING_MESSAGE, ping.Encode())
}

func (ping *PingMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(ping.Nonce)
	return output.Buffer.Bytes()
}

func (ping *PingMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(ping.Header.getBytes()).WriteBytes(ping.GetPayload())
	return output.Buffer.Bytes()
}
