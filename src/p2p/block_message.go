package p2p

import (
	"common"
	"datastruct"
)

type BlockMessage struct {
	Header message_header
	Block  datastruct.Block
}

func (blockMsg *BlockMessage) Decode(payload []byte) {
	blockMsg.Header.init(BLOCK, payload)
	input := common.NewBitcoinInput(payload)
	blockMsg.Block.Init(input)
}

func (block *BlockMessage) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(block.Header.getBytes()).WriteBytes(block.getPayload())
	return output.Buffer.Bytes()
}

func (block *BlockMessage) getPayload() []byte {
	return block.Block.GetBytes()
}
