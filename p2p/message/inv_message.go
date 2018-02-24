package message

import (
	"common"
	"datastruct"
)

//https://bitcoin.org/en/developer-reference#inv
type InvMessage struct {
	header messageHeader

	inventory []datastruct.InvVect
}

func (invMsg *InvMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	invMsg.header.decode(input)
	invCount, _ := input.ReadVarInt()
	invMsg.inventory = make([]datastruct.InvVect, invCount)
	for _, inventory := range invMsg.inventory {
		inventory.Init(input)
	}
}
func (invMsg *InvMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	for _, inventory := range invMsg.inventory {
		output.WriteBytes(inventory.GetBytes())
	}
	return output.Buffer.Bytes()
}

func (invMsg *InvMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(invMsg.header.getBytes()).WriteBytes(invMsg.GetPayload())
	return output.Buffer.Bytes()
}
