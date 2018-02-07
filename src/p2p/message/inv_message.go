package message

import (
	"datastruct"
	"common"
	"constants"
)

//https://bitcoin.org/en/developer-reference#inv
type InvMessage struct {
	Header message_header

	Inventory []datastruct.InvVect
}

func (invMsg *InvMessage) Decode(payload []byte) {
	invMsg.Header.init(constants.INV_MESSAGE, payload)
	input := common.NewBitcoinInput(payload)
	invCount, _ := input.ReadVarInt()
	invMsg.Inventory = make([]datastruct.InvVect, invCount)
	for _, inventory := range invMsg.Inventory {
		inventory.Init(input)
	}
}
func (invMsg *InvMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	for _, inventory := range invMsg.Inventory {
		output.WriteBytes(inventory.GetBytes())
	}
	return output.Buffer.Bytes()
}

func (invMsg *InvMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(invMsg.Header.getBytes()).WriteBytes(invMsg.GetPayload())
	return output.Buffer.Bytes()
}
