package message

import (
	"common"
	"datastruct"
)

type MerkleBlockMessage struct {
	header messageHeader

	//The block header in the format described in the block header section.
	blockHeader datastruct.Header

	//The number of transactions in the block (including ones that don’t match the filter).
	txCount uint32

	//One or more hashes of both transactions and merkle nodes in internal byte order.
	// Each hash is 32 bytes.
	hashes [][]byte

	//A sequence of bits packed eight in a byte with the least significant bit first.
	// May be padded to the nearest byte boundary but must not contain any more bits than that.
	// Used to assign the hashes to particular nodes in the merkle tree as described below.
	flags []byte
}

func (mBlkMsg *MerkleBlockMessage) Decode(contentbytes []byte) {
	input := common.NewBitcoinInput(contentbytes)
	mBlkMsg.header.decode(input)
	mBlkMsg.blockHeader.Init(input)
	input.ReadNum(&mBlkMsg.txCount)
	hashBytesCount, _ := input.ReadVarInt()
	mBlkMsg.hashes = make([][]byte, hashBytesCount)
	for index, _ := range mBlkMsg.hashes {
		mBlkMsg.hashes[index] = make([]byte, 32)
		input.ReadBytes(mBlkMsg.hashes[index])
	}
	flagsBytesCount, _ := input.ReadVarInt()
	mBlkMsg.flags = make([]byte, flagsBytesCount)
	input.ReadBytes(mBlkMsg.flags)
}

func (mBlkMsg *MerkleBlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(mBlkMsg.header.getBytes()).WriteBytes(mBlkMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (mBlkMsg *MerkleBlockMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(mBlkMsg.blockHeader.GetBytes()).WriteNum(mBlkMsg.txCount).
		WriteVarInt(int64(len(mBlkMsg.hashes)))
	for _, hashes := range mBlkMsg.hashes {
		output.WriteBytes(hashes)
	}
	output.WriteVarInt(int64(len(mBlkMsg.flags))).WriteBytes(mBlkMsg.flags)
	return output.Buffer.Bytes()
}
