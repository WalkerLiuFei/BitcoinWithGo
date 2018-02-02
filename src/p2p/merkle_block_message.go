package p2p

import "datastruct"

type MerkleBlockMessage struct {
	//The block header in the format described in the block header section.
	BlockHeader datastruct.Header

	//The number of transactions in the block (including ones that don’t match the filter).
	TXCount uint32

	//One or more hashes of both transactions and merkle nodes in internal byte order.
	// Each hash is 32 bytes.
	Hashes []byte

	//A sequence of bits packed eight in a byte with the least significant bit first.
	// May be padded to the nearest byte boundary but must not contain any more bits than that.
	// Used to assign the hashes to particular nodes in the merkle tree as described below.
	flags []byte
}
