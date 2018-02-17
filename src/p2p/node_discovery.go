package p2p

import (
	"configs"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"p2p/message"
	"persistence"
	"sync"
	"utils"
)

var DNS_SEEDS = []string{"bitseed.xf2.org", "dnsseed.bluematt.me", "seed.bitcoin.sipa.be", "dnsseed.bitcoin.dashjr.org", "seed.bitcoinstats.com"}

type validateCallback func(bool)

type validateAddrArrayCallback func(*SafeNodeValidateCounter)

//安全的并发计数器，计数已经验证过的节点的数量
type SafeNodeValidateCounter struct {
	//lock锁
	lock sync.Mutex
	//已经验证过的节点的数量
	validateCount int
	//有效的地址
	validateAddrs []*net.TCPAddr
}

func (validateCounter *SafeNodeValidateCounter) appendAddr(addr *net.TCPAddr) {
	validateCounter.lock.Lock()
	if validateCounter.validateAddrs == nil {
		validateCounter.validateAddrs = make([]*net.TCPAddr, 0)
	}
	validateCounter.validateAddrs = append(validateCounter.validateAddrs, addr)
	validateCounter.lock.Unlock()
}

func (validateCounter *SafeNodeValidateCounter) increase() {
	validateCounter.lock.Lock()
	validateCounter.validateCount++
	validateCounter.lock.Unlock()
}

// compare with the dst num  == with return  0 , smaller with return  -1 ,greater will return 1
func (validateCounter *SafeNodeValidateCounter) compare(dst int) int {
	validateCounter.lock.Lock()
	if validateCounter.validateCount > dst {
		return 1
	} else if validateCounter.validateCount < dst {
		return -1
	} else {
		return 0
	}
	validateCounter.lock.Unlock()
	return 0
}

//更新本地有效的Node地址
func UpdateUsefulNode() {
	//在直接通过DNS Seed之前应该先查询持久化的Node节点来通信
	tcpAddrChan := make(chan []*net.TCPAddr)
	go ConsumeDNSSeed(tcpAddrChan)
	consumed := 0
	waitValidateAddrArray := make([]*net.TCPAddr, 0)
	for addrArray := range tcpAddrChan {
		consumed++
		waitValidateAddrArray = append(waitValidateAddrArray, addrArray...)
		if consumed == len(DNS_SEEDS) {
			break
		}
	}
	validateAddrArray(waitValidateAddrArray, func(counter *SafeNodeValidateCounter) {
		persistence.StoreUsefulNodes(counter.validateAddrs)
	})
}
func validateAddrArray(addrArray []*net.TCPAddr, callback validateAddrArrayCallback) {
	safeCounter := &SafeNodeValidateCounter{}
	for _, addr := range addrArray {
		if addr == nil || addr.IP == nil {
			//空的地址信息，直接作为已经验证了的
			safeCounter.increase()
			continue
		}
		go ValidateAddr(addr, func(flag bool) {
			safeCounter.increase()
			if flag {
				utils.GetSugarLogger().Infof("Get useful node %s", addr.IP.String())
				safeCounter.appendAddr(addr)
			}
			//FIXME ： 这里只做了一个大约的值限制，应该还有更加精确！
			if safeCounter.compare(len(addrArray)*3/4) >= 0 {
				callback(safeCounter)
			}
		})
	}
}
func ConsumeDNSSeed(tcpAddrChan chan []*net.TCPAddr) {
	result := make([](*net.TCPAddr), len(DNS_SEEDS))
	for _, dnsAddr := range DNS_SEEDS {
		ipArray, err := net.LookupIP(dnsAddr)
		checkerr(err, ConsumeDNSSeed)
		for _, ip := range ipArray {
			tcpAddr := &net.TCPAddr{
				IP:   ip,
				Port: viper.GetInt(configs.PORT),
			}
			result = append(result, tcpAddr)
		}
		tcpAddrChan <- result
	}

}

//链接成功返回true
func ValidateAddr(addr *net.TCPAddr, callback validateCallback) {
	conn, err := net.DialTCP("tcp", nil, addr)
	checkerr(err, ValidateAddr)
	if conn == nil {
		callback(false)
		return
	}
	defer conn.Close()
	//FIXME : 连接是不是必须要关掉？
	checkerr(err, ValidateAddr)
	verMsg := message.InitVersionMessage(viper.GetInt(configs.CURRENT_HEIGHT), addr)
	var versionAckMsg []byte
	_, err = conn.Write(verMsg.Encode())
	checkerr(err, ValidateAddr)
	//等待11S
	for count := 0; count < 11 && len(versionAckMsg) == 0; count++ {
		versionAckMsg, err = ioutil.ReadAll(conn)
	}
	checkerr(err, ValidateAddr)
	msg, err := message.DecodeMessage(versionAckMsg)
	checkerr(err, ValidateAddr)
	if msg == nil {
		callback(false)
		return
	}
	utils.GetSugarLogger().Infof("connect to %s", addr.IP.String())
	callback(true)
}
