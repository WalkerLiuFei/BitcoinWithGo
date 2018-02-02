package p2p

//has no payload message ,like verack ,sendheaders ,getaddr,filterclear ,alter,mempool
type NonePayloadMessage struct {
	Header message_header
}

func (nonePayloadMsg *NonePayloadMessage) Init(msgType MessageType) {
	nonePayloadMsg.Header.init(msgType, []byte{})
}
