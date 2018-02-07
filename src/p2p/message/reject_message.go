package message

import (
	"common"
	"constants"
)

//The reject message informs the receiving node that one of its previous messages has been rejected.
// reference : https://bitcoin.org/en/developer-reference#Reject
type Reject_Message struct {
	Header message_header

	//The type of message rejected as ASCII text without null padding.
	MessageType []byte

	Code byte
	//string type actually
	Reason string

	//string type actually
	ExtraData []byte
}

func (rejectMsg *Reject_Message) Decode(payload []byte) {
	rejectMsg.Header.init(constants.REJECT_MESSAGE, payload)
	input := common.NewBitcoinInput(payload)
	typeBytesCount, err := input.ReadVarInt()
	if err != nil {
		println(err)
		return
	}
	rejectMsg.MessageType = make([]byte, typeBytesCount)
	input.ReadBytes(rejectMsg.MessageType)
	rejectMsg.Code, err = input.InputBuffer.ReadByte()
	if err != nil {
		println(err)
		return
	}
	reasonBytesCount, err := input.ReadVarInt()
	reasonBytes := make([]byte, reasonBytesCount)
	input.ReadBytes(reasonBytes)
	rejectMsg.Reason = string(reasonBytes)
	rejectMsg.ExtraData = input.InputBuffer.Bytes()
}


