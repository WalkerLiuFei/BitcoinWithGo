package datastruct


/**
	区块是比特币系统中最重要的数据结构,其包括区块头和打包的交易
 */
type Block struct {
	Header *Header
	//区块内的交易
	TXns Transaction
}

func NewBlock(transaction []*Transaction, preBlockHash []byte, height int) *Block {

}
