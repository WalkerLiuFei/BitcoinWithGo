package message

import (
	"common"
)

//has no payload message ,like verack ,sendheaders ,getaddr,filterclear ,alter,mempool

type NonePayloadMessage struct {
	Header messageHeader
}

func (nonePayloadMsg *NonePayloadMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	nonePayloadMsg.Header.decode(input)
}
func (nonePayloadMsg *NonePayloadMessage) Encode() []byte {
	return nonePayloadMsg.Header.getBytes()
}
func (nonePayloadMsg *NonePayloadMessage) GetPayload() []byte {
	return []byte{}
}
