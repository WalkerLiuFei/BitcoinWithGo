package message

import (
	"common"
	"constants"
	"math/rand"
)

//https://bitcoin.org/en/developer-reference#ping
type PingMessage struct {
	header messageHeader
	//random number to confirm connection
	nonce uint64
}

func (ping *PingMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	ping.header.decode(input)
	input.ReadNum(&ping.nonce)
}

func (ping *PingMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(ping.nonce)
	return output.Buffer.Bytes()
}

func (ping *PingMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(ping.header.getBytes()).WriteBytes(ping.GetPayload())
	return output.Buffer.Bytes()
}
