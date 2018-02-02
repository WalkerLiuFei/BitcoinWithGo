package p2p

type GetAddressMessage struct {
	Header message_header
}

func (getAddressMessage *GetAddressMessage) Decode([]byte) {
	getAddressMessage.Header.init(GET_ADDRESS, []byte{})
}

func (getAddressMsg *GetAddressMessage) GetBytes() []byte {
	return getAddressMsg.Header.getBytes()
}
