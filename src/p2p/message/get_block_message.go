package message

import (
	"common"
	"constants"
)

//https://bitcoin.org/en/developer-reference#getblocks
//https://bitcoin.org/en/developer-guide#blocks-first
type GetBlockMessage struct {
	Header message_header

	//版本号信息
	Version uint32

	//One or more block header hashes (32 bytes each)
	HeaderHashes [][]byte

	//The header hash of the last header hash being requested;
	// set to zero to get as many blocks as possible
	StopHash []byte
}

//headerHashes :
func (getBlockMsg *GetBlockMessage) Init(hashStop []byte, headerHashes ...[]byte) {
	getBlockMsg.Version = constants.PROTOCOL_VERSION
	getBlockMsg.StopHash = hashStop
	getBlockMsg.HeaderHashes = headerHashes
	getBlockMsg.Header.init(constants.GET_BLOCKS, getBlockMsg.GetPayload())
}
func (getBlockMsg *GetBlockMessage) Decode(payload []byte) {
	getBlockMsg.Header.init(constants.GET_HEADERS, payload)
	input := common.NewBitcoinInput(payload)
	input.ReadNum(&getBlockMsg.Version)
	headerHashCount, err := input.ReadVarInt()
	if err != nil {
		return
	}
	getBlockMsg.HeaderHashes = make([][]byte, headerHashCount)
	for index, _ := range getBlockMsg.HeaderHashes {
		getBlockMsg.HeaderHashes[index] = make([]byte, 32)
		input.ReadBytes(getBlockMsg.HeaderHashes[index])
	}
	getBlockMsg.StopHash = make([]byte, 32)
	input.ReadBytes(getBlockMsg.StopHash)
}

func (getBlockMsg *GetBlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(getBlockMsg.Header.getBytes()).WriteBytes(getBlockMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (getBlockMsg *GetBlockMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(getBlockMsg.Version)
	for _, hash := range getBlockMsg.HeaderHashes {
		output.WriteBytes(hash)
	}
	output.WriteBytes(getBlockMsg.StopHash)
	return output.Buffer.Bytes()
}
