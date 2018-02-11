package constants

type RejectMsgType byte

const (
	//Message could not be decoded.
	//In Reply To any message
	ERROR_DECODE RejectMsgType = 0x01

	//block is invalid for some reason
	//In reply to block message
	BLOCK_INVALID RejectMsgType = 0x10

	TRANSACTION_INVALID RejectMsgType = 0x10
)

//message的类型
type MessageType string

const (
	//版本信息
	VERSION MessageType = "version"

	//收到版本信息后的响应
	VAERSION_ACK MessageType = "verack"
	//区块
	BLOCK MessageType = "block"

	//Get Headers
	GET_HEADERS MessageType = "getheaders"

	//Get Blocks
	GET_BLOCKS MessageType = "getblocks"

	//Get Data
	GET_DATA MessageType = "getdata"

	//Address
	ADDRESS_MESSAGE MessageType = "addr"

	//header
	HEADER MessageType = "header"

	//Get Address
	GET_ADDRESS MessageType = "getaddr"

	//inv message
	INV_MESSAGE MessageType = "inv"

	//ping message
	PING_MESSAGE MessageType = "ping"

	//pong message
	PONG_MESSAGE MessageType = "pong"

	//reject message
	REJECT_MESSAGE MessageType = "reject"

	//mem pool message, reference : https://bitcoin.org/en/developer-reference#mempool
	MEM_POOL_MESSAGE MessageType = "mempool"

	//reference : https://bitcoin.org/en/developer-reference#notfound
	NOT_FOUND MessageType = "notfound"

	//merkle block : https://bitcoin.org/en/developer-reference#merkleblock
	MERKLE_BLOCK MessageType = "merkleblock"
)
