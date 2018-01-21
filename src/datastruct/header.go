package datastruct

import (
	"bytes"
	"common"
)

type Header struct {
	//时间戳的值
	TimeStamp uint32

	//The reference to a Merkle tree collection which is a hash of all transactions related to this block
	MerkleHash []byte

	/**
	 * 猜中这个区块对应的hash难度 使用了的次数
	 */
	Bits uint32

	/**
	 * uint32, The Nonce used to generate this block to allow variations of the
	 * header and compute different hashes
	 */
	Nonce uint32
	//上一个区块的hash值
	PreBlockHash []byte

	//区块的版本号
	Version uint32

	//当前区块的hash值
	BlcokHash []byte
}

func (header *Header) New(input *common.BitcoinInput) {
	input.ReadNum(&header.Version)

	header.MerkleHash = make([]byte, 32)
	input.ReadBytes(header.PreBlockHash)

	header.MerkleHash = make([]byte, 32)
	input.ReadBytes(header.MerkleHash)

	input.ReadNum(&header.TimeStamp)
	input.ReadNum(&header.Bits)
	input.ReadNum(&header.Nonce)
}

func (header *Header) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(header.Version).
		WriteBytes(header.PreBlockHash).
		WriteBytes(header.MerkleHash).
		WriteNum(header.TimeStamp).
		WriteNum(header.Bits).
		WriteNum(header.Nonce)
	return output.Buffer.Bytes()
}
