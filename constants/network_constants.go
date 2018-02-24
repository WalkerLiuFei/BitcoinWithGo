package constants

const (
	//主干网络的端口
	MAIN_NET_PORT uint16 = 8333

	//全网络的测试的端口
	GLOBE_TEST_NET_PORT uint16 = 18333

	//本地网络的端口
	LOCAL_TEST_NET_PORT uint16 = 18444
)

//节点服务的类型
type ServiceType uint64

const (
	// not full node
	UNNAMED ServiceType = 0x00

	//is full nodes
	NODE_NETWORK ServiceType = 0x01
)
