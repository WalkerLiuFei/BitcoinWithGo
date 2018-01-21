package datastruct



const (
	//any data with this num may be ignored
	ERROR = 0

	//	Hash 值对应一个交易
	MSG_TX = 1

	//Hash 值对应一个区块
	MSG_BLOCK = 2

	//Hash of a block header; identical to MSG_BLOCK. Only to be used in
	//getdata message. Indicates the reply should be a merkleblock message
	//rather than a block message; this only works if a bloom filter has been set.

	MSG_FILTERED_BLOCK = 3

	/*
	 * Hash of a block header; identical to MSG_BLOCK. Only to be used in
	 * getdata message. Indicates the reply should be a cmpctblock message. See
	 * BIP 152 for more info.
	 */
	MSG_CMPCT_BLOCK = 4
)
//Inventory vectors are used for notifying other nodes about objects they have or data which is being requested.
//Inventory vectors consist of the following data format:
type InvVect struct {
	//Hash的类型，对应上面的四种
	Type uint32

	//对应 object hash值
	Hash []byte
}
