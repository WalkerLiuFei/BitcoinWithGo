package p2p

import "common"

//getheader message 作用几乎和getblock一样，只不过getheader 只返回header 信息（一次最多2000）

//https://bitcoin.org/en/developer-reference#getheaders

//https://bitcoin.org/en/developer-reference#getblocks
//https://bitcoin.org/en/developer-guide#blocks-first
type GetHeaderMessage struct {
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
func (getBlockMsg *GetHeaderMessage) Init(hashStop []byte, headerHashes ...[]byte) {
	getBlockMsg.Version = PROTOCOL_VERSION
	getBlockMsg.StopHash = hashStop
	getBlockMsg.HeaderHashes = headerHashes
	getBlockMsg.Header.init(GET_BLOCKS, getBlockMsg.GetPayload())
}
func (getBlockMsg *GetHeaderMessage) DirectInit(payload []byte) {
	getBlockMsg.Header.init(GET_HEADERS, payload)
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
	input.ReadBytes(getBlockMsg.StopHash)
}

func (getBlockMsg *GetHeaderMessage) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(getBlockMsg.Header.getBytes()).WriteBytes(getBlockMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (getBlockMsg *GetHeaderMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(getBlockMsg.Version)
	for _, hash := range getBlockMsg.HeaderHashes {
		output.WriteBytes(hash)
	}
	output.WriteBytes(getBlockMsg.StopHash)
	return output.Buffer.Bytes()
}
