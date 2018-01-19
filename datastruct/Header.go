package datastruct

type Header struct {
	//时间戳的值
	TimeStamp uint32

	//The reference to a Merkle tree collection which is a hash of all transactions related to this block
	merkle_root []byte

	/**
	 * 猜中这个区块对应的hash难度 使用了的次数
	 */
	Bits uint32

	/**
	 * uint32, The nonce used to generate this block to allow variations of the
	 * header and compute different hashes
	 */
	nonce uint32
	//上一个区块的hash值
	PreBlockHash []byte

	//区块的版本号
	version uint32

	//当前区块的hash值
	BlcokHash []byte
}

