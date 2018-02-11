package message

import (
	"common"
	"constants"
	"math/rand"
)

//https://bitcoin.org/en/developer-reference#ping
type PongMessage struct {
	header messageHeader
	//random number to confirm connection
	nonce uint64
}

func (pong *PongMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	pong.header.decode(input)
	input.ReadNum(&pong.nonce)
}

func (pong *PongMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(pong.nonce)
	return output.Buffer.Bytes()
}
func (pong *PongMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(pong.header.getBytes()).WriteBytes(pong.GetPayload())
	return output.Buffer.Bytes()
}
