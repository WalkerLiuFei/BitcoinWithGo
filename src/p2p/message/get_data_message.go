package message

import (
	"common"
	"constants"
	"datastruct"
)

type GetDataMessage struct {
	header messageHeader
	//数据及类型
	inventory []datastruct.InvVect
}

func (getDataMsg *GetDataMessage) Init(invType datastruct.InventoryType, hashes ...[]byte) {
	getDataMsg.inventory = make([]datastruct.InvVect, len(hashes))
	for index, _ := range getDataMsg.inventory {
		getDataMsg.inventory[index].Type = invType
		getDataMsg.inventory[index].Hash = hashes[index]
	}
	getDataMsg.header.init(constants.GET_DATA, getDataMsg.GetPayload())
}

func (getDataMsg *GetDataMessage) Decode(contentbytes []byte) {
	input := common.NewBitcoinInput(contentbytes)
	getDataMsg.header.decode(input)
	inventoryCount, _ := input.ReadVarInt()
	getDataMsg.inventory = make([]datastruct.InvVect, inventoryCount)
	for _, inv := range getDataMsg.inventory {
		inv.Init(input)
	}
}
func (getDataMsg *GetDataMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(getDataMsg.header.getBytes()).WriteBytes(getDataMsg.GetPayload())
	return output.Buffer.Bytes()
}
func (getDataMsg *GetDataMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	for _, data := range getDataMsg.inventory {
		output.WriteBytes(data.GetBytes())
	}
	return output.Buffer.Bytes()
}
