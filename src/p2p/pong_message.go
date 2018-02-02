package p2p

import (
	"math/rand"
	"common"
)

//https://bitcoin.org/en/developer-reference#ping
type PongMessage struct {
	Header message_header
	//random number to confirm connection
	Nonce uint64
}

func (pong *PongMessage) Init([]byte) {
	pong.Nonce = uint64(rand.Int63())
	pong.Header.init(PONG_MESSAGE, pong.GetBytes())
}

func (pong *PongMessage) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(pong.Nonce)
	return output.Buffer.Bytes()
}
