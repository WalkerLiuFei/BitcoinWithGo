package message

import (
	"common"
	"datastruct"
	"constants"
)

type BlockMessage struct {
	Header message_header
	Block  datastruct.Block
}

func (blockMsg *BlockMessage) Decode(payload []byte) {
	blockMsg.Header.init(constants.BLOCK, payload)
	input := common.NewBitcoinInput(payload)
	blockMsg.Block.Init(input)
}

func (block *BlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(block.Header.getBytes()).WriteBytes(block.GetPayload())
	return output.Buffer.Bytes()
}

func (block *BlockMessage) GetPayload() []byte {
	return block.Block.GetBytes()
}
