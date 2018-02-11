package message

import (
	"common"
	"constants"
	"datastruct"
)

type HeaderMessage struct {
	//header
	header messageHeader
	//所有的区块头
	blockHeaders []datastruct.Header
}

func (headerMsg *HeaderMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	headerMsg.header.decode(input)
	headerCount, err := input.ReadVarInt()
	checkError(err)
	headerMsg.blockHeaders = make([]datastruct.Header, headerCount)
	for index, _ := range headerMsg.blockHeaders {
		headerMsg.blockHeaders[index].Init(input)
		//ignore tx count , always 0X00
		input.ReadByte()
	}
}

func (headerMsg *HeaderMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(headerMsg.GetPayload()).WriteBytes(headerMsg.GetPayload())
	return output.Buffer.Bytes()
}
func (headerMsg *HeaderMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteVarInt(int64(len(headerMsg.blockHeaders)))
	for _, header := range headerMsg.blockHeaders {
		output.WriteBytes(header.GetBytes())
	}
	return output.Buffer.Bytes()
}
