package constants

const PROTOCOL_VERSION = 70014

//
const USER_AGENT = "/Satoshi:0.7.2/"

type NetContext struct {
	//网络
	Network string
	//端口号
	DefaultPort uint16
	//Start Bytes
	StartBytes uint32
	//Max nBits -->最低难度
	MaxNBits uint32
}

var /*constant*/ (
	MAIN_NET = NetContext{
		Network:     "Mainnet",
		DefaultPort: 8333,
		StartBytes:  0xf9beb4d9,
		MaxNBits:    0X1D00FFFF,
	}

	TEST_NET = NetContext{
		Network:     "Testnet",
		DefaultPort: 18333,
		StartBytes:  0x0b110907,
		MaxNBits:    0x1d00ffff,
	}

	REG_TEST_NET = NetContext{
		Network:     "Regtest",
		DefaultPort: 18444,
		StartBytes:  0xfabfb5da,
		MaxNBits:    0x207fffff,
	}
)
