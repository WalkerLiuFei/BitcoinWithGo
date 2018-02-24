package datastruct

import (
	"common"
	"constants"
	"net"
	"time"
	"utils"
)

//当某个地方需要网络地址时，就使用这个结构。网络地址在版本信息中没有以时间戳作为前缀
type NetworkAddress struct {
	// the Time (version >= 31402). Not present in version message.
	// reference : https://bitcoin.org/en/developer-reference#protocol-versions
	Time uint32

	//	same service(s) listed in version
	Services uint64

	//IPv6 address. Network byte order. The original client only supported IPv4 and only read the
	// last 4 bytes to get the IPv4 address. However, the IPv4 address is written into the message as a
	// 16 byte IPv4-mapped IPv6 address (12 bytes 00 00 00 00 00 00 00 00 00 00 FF FF, followed by the 4 bytes of the IPv4 address).
	IP net.IP

	//uint16 port number
	Port uint16
}

/**
	Hexdump example of Network address structure

0000   01 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  ................
0010   00 00 FF FF 0A 00 00 01  20 8D                    ........ .

 Network address:
 01 00 00 00 00 00 00 00                         - 1 (NODE_NETWORK: see services listed under version command)
 00 00 00 00 00 00 00 00 00 00 FF FF 0A 00 00 01 - IPv6: ::ffff:a00:1 or IPv4: 10.0.0.1
 20 8D                                           - Port 8333
*/

func (networkAddr *NetworkAddress) InitNetWorkAddress(input common.BitcoinInput) {
	if constants.PROTOCOL_VERSION < 31402 {
		err := input.ReadNum(&networkAddr.Time)
		if err != nil {
			return
		}
	}
	input.ReadNum(&networkAddr.Services)
	networkAddr.IP = make([]byte, 16)
	input.ReadBytes(networkAddr.IP)
	input.ReadNum(&networkAddr.Port)
}

func InitNetWorkAddress(input common.BitcoinInput) *NetworkAddress {
	networkAddr := &NetworkAddress{}
	if constants.PROTOCOL_VERSION < 31402 {
		err := input.ReadNum(&networkAddr.Time)
		checkerr(err, InitNetWorkAddress)
	}
	input.ReadNum(&networkAddr.Services)
	networkAddr.IP = make([]byte, 16)
	input.ReadBytes(networkAddr.IP)
	input.ReadNum(&networkAddr.Port)
	return networkAddr
}
func InitByTCPAddr(addr *net.TCPAddr, serviceType constants.ServiceType) *NetworkAddress {
	netWorkAddress := &NetworkAddress{
		Services: uint64(serviceType),
		IP:       addr.IP,
		Port:     uint16(addr.Port),
	}
	//FIXME : 这个时间戳信息到底要不要？
	//if constants.PROTOCOL_VERSION < 31402 {
	if true {
		netWorkAddress.Time = uint32(time.Now().Unix())
	}
	return netWorkAddress
}
func (networkAddr *NetworkAddress) Encode() []byte {
	output := common.BitcoinOuput{}
	if constants.PROTOCOL_VERSION < 31402 {
		output.WriteNum(networkAddr.Time)
	}
	output.WriteNum(networkAddr.Services).
		WriteBytes(networkAddr.IP).
		WriteNum(networkAddr.Port)
	return output.Buffer.Bytes()
}
func checkerr(e error, funcName interface{}) {
	logger := utils.GetSugarLogger()
	if e != nil {
		logger.Errorf("call func name is %s,msg is %s", utils.GetFunctionName(funcName), e.Error())
	}
}