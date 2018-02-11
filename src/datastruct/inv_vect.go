package datastruct

import "common"

type InventoryType uint32

const (
	//any data with this num may be ignored
	ERROR InventoryType = 0

	//	Hash 值对应一个交易
	MSG_TX InventoryType = 1

	//Hash 值对应一个区块
	MSG_BLOCK InventoryType = 2

	//Hash of a block header; identical to MSG_BLOCK. Only to be used in
	//getdata message. Indicates the reply should be a merkleblock message
	//rather than a block message; this only works if a bloom filter has been set.

	MSG_FILTERED_BLOCK InventoryType = 3

	/*
	 * Hash of a block header; identical to MSG_BLOCK. Only to be used in
	 * getdata message. Indicates the reply should be a cmpctblock message. See
	 * BIP 152 for more info.
	 */
	MSG_CMPCT_BLOCK InventoryType = 4
)

//inventory vectors are used for notifying other nodes about objects they have or data which is being requested.
//inventory vectors consist of the following data format:
//used to identify transactions and blocks available for download through the Bitcoin P2P network.
//一般在P2P网络中，作为 交易或者区块的验证
type InvVect struct {
	//Hash的类型，对应上面的四种
	Type InventoryType

	//block header / transaction 的Hash 值
	Hash []byte
}

func (invVect *InvVect) Init(input common.BitcoinInput) {
	input.ReadNum(&invVect.Type)
	invVect.Hash = make([]byte, 32)
	input.ReadBytes(invVect.Hash)
}

func (invVect *InvVect) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(invVect.Type).WriteBytes(invVect.Hash)
	return output.Buffer.Bytes()
}
