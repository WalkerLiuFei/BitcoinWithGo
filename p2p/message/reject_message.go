package message

import (
	"common"
)

//The reject message informs the receiving node that one of its previous messages has been rejected.
// reference : https://bitcoin.org/en/developer-reference#Reject
type Reject_Message struct {
	header messageHeader

	//The type of message rejected as ASCII text without null padding.
	messageType []byte

	code byte
	//string type actually
	reason string

	//string type actually
	extraData []byte
}

func (rejectMsg *Reject_Message) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(rejectMsg.header.getBytes()).WriteBytes(rejectMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (rejectMsg *Reject_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(rejectMsg.messageType).
		WriteBytes([]byte(rejectMsg.reason)).
		WriteBytes(rejectMsg.extraData)
	return output.Buffer.Bytes()
}
func (rejectMsg *Reject_Message) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	rejectMsg.header.decode(input)
	typeBytesCount, err := input.ReadVarInt()
	if err != nil {
		println(err)
		return
	}
	rejectMsg.messageType = make([]byte, typeBytesCount)
	input.ReadBytes(rejectMsg.messageType)
	rejectMsg.code, err = input.InputBuffer.ReadByte()
	if err != nil {
		println(err)
		return
	}
	reasonBytesCount, err := input.ReadVarInt()
	reasonBytes := make([]byte, reasonBytesCount)
	input.ReadBytes(reasonBytes)
	rejectMsg.reason = string(reasonBytes)
	rejectMsg.extraData = input.InputBuffer.Bytes()
}
