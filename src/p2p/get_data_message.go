package p2p

import (
	"common"
	"datastruct"
)

type GetDataMessage struct {
	Header message_header
	//数据及类型
	Inventory []datastruct.InvVect
}

func (getDataMsg *GetDataMessage) Encode(invType datastruct.InventoryType, hashes ...[]byte) {
	getDataMsg.Inventory = make([]datastruct.InvVect, len(hashes))
	for index, _ := range getDataMsg.Inventory {
		getDataMsg.Inventory[index].Type = invType
		getDataMsg.Inventory[index].Hash = hashes[index]
	}
	getDataMsg.Header.init(GET_DATA, getDataMsg.GetPayload())
}

func (getDataMsg *GetDataMessage) Decode(payload []byte) {
	input := common.NewBitcoinInput(payload)
	getDataMsg.Header.init(GET_DATA, payload)
	inventoryCount, _ := input.ReadVarInt()
	getDataMsg.Inventory = make([]datastruct.InvVect, inventoryCount)
	for _, inv := range getDataMsg.Inventory {
		inv.Init(input)
	}
}

func (getDataMsg *GetDataMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	for _, data := range getDataMsg.Inventory {
		output.WriteBytes(data.GetBytes())
	}
	return output.Buffer.Bytes()
}
