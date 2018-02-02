package p2p

const PROTOCOL_VERSION = 70014

type PortType uint16

const (
	//主干网络的端口
	MAIN_NET_PORT PortType = 8333

	//全网络的测试的端口
	GLOBE_TEST_NET_PORT PortType = 18333

	//本地网络的端口
	LOCAL_TEST_NET_PORT PortType = 18444
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

