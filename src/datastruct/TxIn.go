package datastruct

type TxIn struct {
	/**
	32 字节,对应交易信息的hash值
	*/
	hash []byte

	/*
	 *	32字节的无符号数，这笔交易在 输入 的位置，默认从零开始
	 */
	index int64

	//签名
	sigScript []byte

	//uint32, Transaction version as defined by the sender. Intended for
	//"replacement" of transactions when information is updated before
	// inclusion into a block.

	sequence uint32
}
