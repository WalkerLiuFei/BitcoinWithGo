package datastruct

type Block struct {
	TimeStamp int64
	//
	Data []byte
	//上一个区块的hash值
	PreBlockHash []byte
	//区块内的交易
	TXns Transaction
}
