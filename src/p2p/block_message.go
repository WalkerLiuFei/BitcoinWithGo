package p2p

import (
	"datastruct"
	"common"
)

type BlockMessage struct {
	Header message_header
	Block  datastruct.Block
}

func (blockMsg *BlockMessage) Init(payload []byte) {
	blockMsg.Header.init(BLOCK, payload)
	input := common.NewBitcoinInput(payload)
	blockMsg.Block.Init(*input)
}
