package message

import (
	"common"
	"constants"
	"datastruct"
)

//https://bitcoin.org/en/developer-reference#addr
type Address_Message struct {
	header *messageHeader

	addressList []datastruct.NetworkAddress
}

func (addrMessage *Address_Message) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	addrMessage.header.decode(input)
	addrCount, _ := input.ReadVarInt()
	addrMessage.addressList = make([]datastruct.NetworkAddress, addrCount)
	for _, networkaddr := range addrMessage.addressList {
		networkaddr.Init(input)
	}
}

func (addrMessage *Address_Message) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(addrMessage.header.getBytes()).WriteBytes(addrMessage.GetPayload())
	return output.Buffer.Bytes()
}

func (addrMessage *Address_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteVarInt(int64(len(addrMessage.addressList)))
	for _, networkaddr := range addrMessage.addressList {
		output.WriteBytes(networkaddr.Encode())
	}
	return output.Buffer.Bytes()
}
