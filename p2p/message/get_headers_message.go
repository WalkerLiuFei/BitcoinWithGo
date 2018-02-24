package message

import (
	"common"
	"constants"
)

//getheader message 作用几乎和getblock一样，只不过getheader 只返回header 信息（一次最多2000）

//https://bitcoin.org/en/developer-reference#getheaders

//https://bitcoin.org/en/developer-reference#getblocks
//https://bitcoin.org/en/developer-guide#blocks-first
type GetHeaderMessage struct {
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
func (getBlockMsg *GetHeaderMessage) Init(hashStop []byte, headerHashes ...[]byte) {
	getBlockMsg.version = constants.PROTOCOL_VERSION
	getBlockMsg.stopHash = hashStop
	getBlockMsg.headerHashes = headerHashes
	getBlockMsg.header.init(constants.GET_BLOCKS, getBlockMsg.GetPayload())
}
func (getBlockMsg *GetHeaderMessage) Decode(payload []byte) {
	input := common.NewBitcoinInput(payload)
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
	input.ReadBytes(getBlockMsg.stopHash)
}

func (getBlockMsg *GetHeaderMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(getBlockMsg.header.getBytes()).WriteBytes(getBlockMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (getBlockMsg *GetHeaderMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(getBlockMsg.version)
	for _, hash := range getBlockMsg.headerHashes {
		output.WriteBytes(hash)
	}
	output.WriteBytes(getBlockMsg.stopHash)
	return output.Buffer.Bytes()
}
