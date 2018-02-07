package message

import "constants"

//has no payload message ,like verack ,sendheaders ,getaddr,filterclear ,alter,mempool

type NonePayloadMessage struct {
	Header message_header
}

var msgType constants.MessageType

func GetMessageType() constants.MessageType {
	return msgType
}

func SetMessageType(messageType constants.MessageType) {
	msgType = messageType
}

func (nonePayloadMsg *NonePayloadMessage) Decode([]byte) {
	nonePayloadMsg.Header.init(msgType, []byte{})
}
func (nonePayloadMsg *NonePayloadMessage) Encode() []byte {
	return nonePayloadMsg.Header.getBytes()
}
func (nonePayloadMsg *NonePayloadMessage) GetPayload() []byte {
	return []byte{}
}
