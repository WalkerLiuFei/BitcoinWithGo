package p2p

type RejectMsgType byte

const (
	//Message could not be decoded.
	//In Reply To any message
	ERROR_DECODE RejectMsgType = 0x01

	//Block is invalid for some reason
	//In reply to block message
	BLOCK_INVALID RejectMsgType = 0x10

	TRANSACTION_INVALID RejectMsgType = 0x10
)
