package message

import (
	"common"
	"datastruct"
)

type BlockMessage struct {
	header *messageHeader
	block  *datastruct.Block
}

func (blockMsg *BlockMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	blockMsg.header.decode(input)
	blockMsg.block.Init(input)
}

func (block *BlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(block.header.getBytes()).WriteBytes(block.GetPayload())
	return output.Buffer.Bytes()
}

func (block *BlockMessage) GetPayload() []byte {
	return block.block.GetBytes()
}
