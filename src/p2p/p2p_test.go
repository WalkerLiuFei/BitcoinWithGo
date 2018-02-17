package p2p

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
	"configs"
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

func TestUpdateUsefulNode(t *testing.T) {
	configs.InitConfigs()
	UpdateUsefulNode()
	for true {
		time.Sleep(time.Second)
	}
}
func TestValidateAddr(t *testing.T) {
	configs.InitConfigs()
	addr := &net.TCPAddr{
		IP:   net.ParseIP("101.201.142.252"),
		Port: 8333,
	}
	ValidateAddr(addr, func(b bool) {
		fmt.Println(b)
	})
}
func TestConnectPeer(t *testing.T) {

	//Test Des
	/*	configs.InitConfigs()
		//addrArray := ConsumeDNSSeed(constants.MAIN_NET_PORT)
		addrArray := ConsumeDNSSeed()
		var conn *PeerConn
		var flag bool
		for _, addr := range addrArray {
			if addr == nil || addr.IP == nil {
				continue
			}
			conn, flag = ValidateAddr(addr)
			if flag {
				break
			}
		}
		headers := GetHeaderMessage(conn)
		for headerMsg := range headers {
			fmt.Println(headerMsg)
		}*/
}

func TestChannel(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	//don't forget the last "()"
	go func() {
		for count := 0; count < 20; count++ {
			time.Sleep(1 * time.Second)
			c1 <- strconv.Itoa(count) + "one"
		}
	}()

	go func() {
		for count := 0; count < 10; count++ {
			time.Sleep(2 * time.Second)
			c2 <- strconv.Itoa(count) + "two"
		}
	}()
	for true {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
