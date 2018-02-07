package message

import (
	"common"
	"datastruct"
	"constants"
)

//https://bitcoin.org/en/developer-reference#addr
type Address_Message struct {
	Header message_header

	AddressList []datastruct.NetworkAddress
}

func (addrMessage *Address_Message) Decode(payload []byte) {
	addrMessage.Header.init(constants.ADDRESS_MESSAGE, payload)
	input := common.NewBitcoinInput(payload)
	addrCount, _ := input.ReadVarInt()
	addrMessage.AddressList = make([]datastruct.NetworkAddress, addrCount)
	for _, networkaddr := range addrMessage.AddressList {
		networkaddr.Init(input)
	}
}

func (addrMessage *Address_Message) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(addrMessage.Header.getBytes()).WriteBytes(addrMessage.GetPayload())
	return output.Buffer.Bytes()
}

func (addrMessage *Address_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteVarInt(int64(len(addrMessage.AddressList)))
	for _, networkaddr := range addrMessage.AddressList {
		output.WriteBytes(networkaddr.Encode())
	}
	return output.Buffer.Bytes()
}
