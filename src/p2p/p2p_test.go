package p2p

import (
	"configs"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestConsumeDNSSeed(t *testing.T) {
	//ipArray := ConsumeDNSSeed(constants.GLOBE_TEST_NET_PORT)
	//fmt.Println(ipArray)
}
func TestConnectPeer2(t *testing.T) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "baidu.com")
	fmt.Println(tcpAddr)
}
func TestUnixSecond(t *testing.T) {
	fmt.Println(time.Now().Unix())
}

func TestConnectPeer(t *testing.T) {
	configs.InitConfigs()
	//addrArray := ConsumeDNSSeed(constants.MAIN_NET_PORT)
	ip := net.IP{}
	ip.UnmarshalText([]byte("62.116.188.85"))
	addr := &net.TCPAddr{
		IP:   ip,
		Port: 8333,
	}
	ConnectPeer(addr)
	/*for _,addr := range addrArray {
		if addr != nil && len(addr.IP) > 0{

		}
	}*/

}
