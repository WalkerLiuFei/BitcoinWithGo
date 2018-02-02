package p2p

import (
	"datastruct"
	"common"
)

//https://bitcoin.org/en/developer-reference#inv
type InvMessage struct {
	Header message_header

	Inventory []datastruct.InvVect
}

func (invMsg *InvMessage) Init(payload []byte) {
	invMsg.Header.init(INV_MESSAGE, payload)
	input := common.NewBitcoinInput(payload)
	invCount, _ := input.ReadVarInt()
	invMsg.Inventory = make([]datastruct.InvVect, invCount)
	for _, inventory := range invMsg.Inventory {
		inventory.Init(input)
	}
}

func (invMsg *InvMessage) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(invMsg.Header.getBytes())
	for _, inventory := range invMsg.Inventory {
		output.WriteBytes(inventory.GetBytes())
	}
	return output.Buffer.Bytes()
}
