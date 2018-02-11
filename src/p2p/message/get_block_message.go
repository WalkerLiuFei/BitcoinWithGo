package message

import (
	"common"
	"constants"
)

//https://bitcoin.org/en/developer-reference#getblocks
//https://bitcoin.org/en/developer-guide#blocks-first
type GetBlockMessage struct {
	header messageHeader

	//版本号信息
	version uint32

	//One or more block header hashes (32 bytes each)
	headerHashes [][]byte

	//The header hash of the last header hash being requested;
	// set to zero to get as many blocks as possible
	stopHash []byte
}

//headerHashes :
func (getBlockMsg *GetBlockMessage) Init(hashStop []byte, headerHashes ...[]byte) {
	getBlockMsg.version = constants.PROTOCOL_VERSION
	getBlockMsg.stopHash = hashStop
	getBlockMsg.headerHashes = headerHashes
	getBlockMsg.header.init(constants.GET_BLOCKS, getBlockMsg.GetPayload())
}
func (getBlockMsg *GetBlockMessage) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	getBlockMsg.header.decode(input)
	input.ReadNum(&getBlockMsg.version)
	headerHashCount, err := input.ReadVarInt()
	if err != nil {
		return
	}
	getBlockMsg.headerHashes = make([][]byte, headerHashCount)
	for index, _ := range getBlockMsg.headerHashes {
		getBlockMsg.headerHashes[index] = make([]byte, 32)
		input.ReadBytes(getBlockMsg.headerHashes[index])
	}
	getBlockMsg.stopHash = make([]byte, 32)
	input.ReadBytes(getBlockMsg.stopHash)
}

func (getBlockMsg *GetBlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(getBlockMsg.header.getBytes()).WriteBytes(getBlockMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (getBlockMsg *GetBlockMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(getBlockMsg.version)
	for _, hash := range getBlockMsg.headerHashes {
		output.WriteBytes(hash)
	}
	output.WriteBytes(getBlockMsg.stopHash)
	return output.Buffer.Bytes()
}
