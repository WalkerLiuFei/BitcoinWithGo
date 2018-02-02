package p2p

import (
	"common"
)

//https://bitcoin.org/en/developer-reference#addr
type Address_Message struct {
	Header message_header

	AddressList []NetworkAddress
}

func (addrMessage *Address_Message) Decode(payload []byte) {
	addrMessage.Header.init(ADDRESS_MESSAGE, payload)
	input := common.NewBitcoinInput(payload)
	addrCount, _ := input.ReadVarInt()
	addrMessage.AddressList = make([]NetworkAddress, addrCount)
	for _, networkaddr := range addrMessage.AddressList {
		networkaddr.Init(input, false)
	}
}

func (addrMessage *Address_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteVarInt(int64(len(addrMessage.AddressList)))
	for _, networkaddr := range addrMessage.AddressList {
		output.WriteBytes(networkaddr.GetBytes(false))
	}
	return output.Buffer.Bytes()
}
