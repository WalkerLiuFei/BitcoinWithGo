package constants

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


//message的类型
type MessageType string

const (
	//版本信息
	VERSION MessageType = "version"

	//区块
	BLOCK MessageType = "block"

	//Get Headers
	GET_HEADERS MessageType = "getheaders"

	//Get Blocks
	GET_BLOCKS MessageType = "getblocks"

	//Get Data
	GET_DATA MessageType = "getdata"

	//Address
	ADDRESS_MESSAGE = "addr"
	//Header
	HEADER = "header"
	//Get Address
	GET_ADDRESS = "getaddr"

	//inv message
	INV_MESSAGE = "inv"

	//ping message
	PING_MESSAGE = "ping"

	//pong message
	PONG_MESSAGE = "pong"

	//reject message
	REJECT_MESSAGE = "reject"

	//mem pool message, reference : https://bitcoin.org/en/developer-reference#mempool
	MEM_POOL_MESSAGE = "mempool"

	//reference : https://bitcoin.org/en/developer-reference#notfound
	NOT_FOUND = "notfound"

	//merkle block : https://bitcoin.org/en/developer-reference#merkleblock
	MERKLE_BLOCK = "merkleblock"
)

