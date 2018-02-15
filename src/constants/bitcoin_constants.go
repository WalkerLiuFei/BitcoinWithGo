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
	Magic uint32
	//Max nBits -->最低难度
	MaxNBits uint32
}

var /*constant*/ (
	MAIN_NET = NetContext{
		Network:     "Mainnet",
		DefaultPort: 8333,
		Magic:       0xf9beb4d9,
		MaxNBits:    0X1D00FFFF,
	}

	TEST_NET = NetContext{
		Network:     "Testnet",
		DefaultPort: 18333,
		Magic:       0x0b110907,
		MaxNBits:    0x1d00ffff,
	}

	REG_TEST_NET = NetContext{
		Network:     "Regtest",
		DefaultPort: 18444,
		Magic:       0xfabfb5da,
		MaxNBits:    0x207fffff,
	}
)

const GENESIS_BLOCK_HASH = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"