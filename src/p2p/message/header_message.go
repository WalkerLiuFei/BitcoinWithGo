package message

import (
	"common"
	"constants"
	"datastruct"
)

type HeaderMessage struct {
	//Header
	Header message_header
	//所有的区块头
	BlockHeaders []datastruct.Header
}

func (headerMsg *HeaderMessage) Decode(payload []byte) {
	headerMsg.Header.init(constants.HEADER, payload)
	input := common.NewBitcoinInput(payload)
	headerCount, err := input.ReadVarInt()
	checkError(err)
	headerMsg.BlockHeaders = make([]datastruct.Header, headerCount)
	for index, _ := range headerMsg.BlockHeaders {
		headerMsg.BlockHeaders[index].Init(input)
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
	output.WriteVarInt(int64(len(headerMsg.BlockHeaders)))
	for _, header := range headerMsg.BlockHeaders {
		output.WriteBytes(header.GetBytes())
	}
	return output.Buffer.Bytes()
}
