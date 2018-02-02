package p2p

import (
	"math/rand"
	"common"
)

//https://bitcoin.org/en/developer-reference#ping
type PingMessage struct {
	Header message_header
	//random number to confirm connection
	Nonce uint64
}

func (ping *PingMessage) Init([]byte) {
	ping.Nonce = uint64(rand.Int63())
	ping.Header.init(PING_MESSAGE, ping.GetBytes())
}

func (ping *PingMessage) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(ping.Nonce)
	return output.Buffer.Bytes()
}
