package datastruct

import (
	"common"
	"utils"
)

/**
区块是比特币系统中最重要的数据结构,其包括区块头和打包的交易
参考 ： https://bitcoin.org/en/developer-reference#block
*/

type Block struct {
	//区块头
	Header *Header
	//区块内的交易
	TXns []Transaction
}

func (block *Block) Init(input common.BitcoinInput) {
	block.Header = &Header{}
	block.Header.Init(input)
	txCount, err := input.ReadVarInt()
	if err != nil {
		logger.Error(err.Error())
	}
	block.TXns = make([]Transaction, txCount)
	for _, transaction := range block.TXns {
		transaction.Init(input)
	}
}

func (block *Block) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(block.Header.GetBytes()).WriteVarInt(int64(len(block.TXns)))
	for _, transaction := range block.TXns {
		output.WriteBytes(transaction.GetBytes())
	}
	return output.Buffer.Bytes()
}

func (block *Block) calculateMerkleHash() []byte {
	hashes := make([][]byte, len(block.TXns))
	for index, tx := range block.TXns {
		hashes[index] = tx.GetTxHash()
	}
	for len(hashes) > 1 {
		hashes = utils.MerkleHash(hashes)
	}
	return hashes[0]
}
