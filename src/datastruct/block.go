package datastruct

import (
	"common"
	"go.uber.org/zap"
)

/**
	区块是比特币系统中最重要的数据结构,其包括区块头和打包的交易
 */

type Block struct {
	//区块头
	Header *Header
	//区块内的交易
	TXns []Transaction
}

func NewBlock(input common.BitcoinInput) *Block {
	header := Header{}
	header.New(&input)
	block := Block{
		Header: &header,
	}
	txCount, err := input.ReadVarInt()
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

}
