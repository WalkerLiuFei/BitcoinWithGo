package message

import "constants"

type GetAddressMessage struct {
	Header message_header
}

func (getAddressMessage *GetAddressMessage) Decode([]byte) {
	getAddressMessage.Header.init(constants.GET_ADDRESS, []byte{})
}

func (getAddressMsg *GetAddressMessage) GetPayload() []byte {
	return []byte{}
}

func (getAddressMsg *GetAddressMessage) Encode() []byte {
	return getAddressMsg.Header.getBytes()
}
