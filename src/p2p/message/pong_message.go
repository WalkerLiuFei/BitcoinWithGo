package message

import (
	"math/rand"
	"common"
	"constants"
)

//https://bitcoin.org/en/developer-reference#ping
type PongMessage struct {
	Header message_header
	//random number to confirm connection
	Nonce uint64
}

func (pong *PongMessage) Decode([]byte) {
	pong.Nonce = uint64(rand.Int63())
	pong.Header.init(constants.PONG_MESSAGE, pong.Encode())
}

func (pong *PongMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(pong.Nonce)
	return output.Buffer.Bytes()
}
func (pong *PongMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(pong.Header.getBytes()).WriteBytes(pong.GetPayload())
	return output.Buffer.Bytes()
}
