package message

import (
	"common"
	"datastruct"
	"constants"
)

type MerkleBlockMessage struct {
	Header message_header

	//The block header in the format described in the block header section.
	BlockHeader datastruct.Header

	//The number of transactions in the block (including ones that don’t match the filter).
	TXCount uint32

	//One or more hashes of both transactions and merkle nodes in internal byte order.
	// Each hash is 32 bytes.
	Hashes [][]byte

	//A sequence of bits packed eight in a byte with the least significant bit first.
	// May be padded to the nearest byte boundary but must not contain any more bits than that.
	// Used to assign the hashes to particular nodes in the merkle tree as described below.
	Flags []byte
}

func (mBlkMsg *MerkleBlockMessage) Decode(payload []byte) {
	mBlkMsg.Header.init(constants.MERKLE_BLOCK, payload)
	input := common.NewBitcoinInput(payload)
	mBlkMsg.BlockHeader.Init(input)
	input.ReadNum(&mBlkMsg.TXCount)
	hashBytesCount, _ := input.ReadVarInt()
	mBlkMsg.Hashes = make([][]byte, hashBytesCount)
	for index, _ := range mBlkMsg.Hashes {
		mBlkMsg.Hashes[index] = make([]byte, 32)
		input.ReadBytes(mBlkMsg.Hashes[index])
	}
	flagsBytesCount, _ := input.ReadVarInt()
	mBlkMsg.Flags = make([]byte, flagsBytesCount)
	input.ReadBytes(mBlkMsg.Flags)
}

func (mBlkMsg *MerkleBlockMessage) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(mBlkMsg.Header.getBytes()).WriteBytes(mBlkMsg.GetPayload())
	return output.Buffer.Bytes()
}

func (mBlkMsg *MerkleBlockMessage) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(mBlkMsg.BlockHeader.GetBytes()).WriteNum(mBlkMsg.TXCount).
		WriteVarInt(int64(len(mBlkMsg.Hashes)))
	for _, hashes := range mBlkMsg.Hashes {
		output.WriteBytes(hashes)
	}
	output.WriteVarInt(int64(len(mBlkMsg.Flags))).WriteBytes(mBlkMsg.Flags)
	return output.Buffer.Bytes()
}
