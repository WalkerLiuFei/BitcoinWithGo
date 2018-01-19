package datastruct

type Transaction struct {
	//int32_t, transaction data format version (signed)
	version uint32

	//a list of 1 or more transaction inputs or sources for coins
	TxIns []TxIn

	// a list of 1 or more transaction outputs or destinations for coins
	TxOuts []TxOut

	/*
	 表示这笔交易被解锁时，指定的区块高度 / 时间戳。
	 当值为0 时，表明这笔交易没有被锁
	 当值 < 500000000 时，表明这个值指的是区块高度
	 当值 > 500000000时，表明这个值是UNIX 时间戳
	 */
	LockTime uint32

	//当前这笔交易的Hash 值
	TxHash []byte
}
