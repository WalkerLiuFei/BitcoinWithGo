package p2p

import (
	"configs"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"p2p/message"
	"sync"
	"utils"
	"math/rand"
	"time"
)

var DNS_SEEDS = []string{"bitseed.xf2.org", "dnsseed.bluematt.me", "seed.bitcoin.sipa.be", "dnsseed.bitcoin.dashjr.org", "seed.bitcoinstats.com"}

type validateCallback func(*net.TCPAddr, bool)

//写的时候加锁，读的时候共享锁
type UsefulNodes struct {
	mutex sync.RWMutex
	nodes []*net.TCPAddr
}

var UsefulNodesHolder *UsefulNodes
//移除失效的,重复的节点
func (usefulNodes *UsefulNodes) FoundAndRemove(addr *net.TCPAddr) {
	if usefulNodes == nil {
		return
	}
	usefulNodes.mutex.Lock()
	for index, nodeAddr := range usefulNodes.nodes {
		if nodeAddr.IP.Equal(addr.IP) {
			if index > 0 && index < len(usefulNodes.nodes)-1 {
				usefulNodes.nodes = append(usefulNodes.nodes[:index-1], usefulNodes.nodes[index+1:]...)
			} else if index > 0 {
				usefulNodes.nodes = usefulNodes.nodes[:index-1]
			} else {
				usefulNodes.nodes = usefulNodes.nodes[index+1:]
			}
		}
	}
	usefulNodes.mutex.Unlock()
}

func (usefulNodes *UsefulNodes) Contain(addr *net.TCPAddr) bool {
	if usefulNodes == nil {
		return false
	}
	for _, nodeAddr := range usefulNodes.nodes {
		if nodeAddr.IP.Equal(addr.IP) {
			return true
		}
	}
	return false
}

func (usefulNodes *UsefulNodes) Append(addr *net.TCPAddr) {
	if usefulNodes == nil {
		return
	}
	usefulNodes.mutex.Lock()
	usefulNodes.nodes = append(usefulNodes.nodes, addr)
	usefulNodes.mutex.Unlock()
}

//随机从缓存中返回一个节点
func (usefulNodes *UsefulNodes) GetRandomNodeSync() *net.TCPAddr {
	if usefulNodes == nil {
		return nil
	}
	index := rand.Intn(len(usefulNodes.nodes))
	return usefulNodes.nodes[index]
}

//线程阻塞的，获取一个有效的地址,如果没有，则阻塞线程等待
func GetRandomNodeAsync() *net.TCPAddr {
	for true {
		addr := UsefulNodesHolder.GetRandomNodeSync()
		if addr != nil {
			return addr
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

//更新本地有效的Node地址
func UpdateUsefulNode() {
	//如果缓存里面有有效的节点，直接更新验证
	if UsefulNodesHolder != nil && len(UsefulNodesHolder.nodes) > 0 {
		validateAddrArray(UsefulNodesHolder.nodes)
		return
	}
	//在直接通过DNS Seed之前应该先查询持久化的Node节点来通信
	tcpAddrChan := make(chan []*net.TCPAddr)
	go ConsumeDNSSeed(tcpAddrChan)
	waitValidateAddrArray := getValidateAddrArray(tcpAddrChan)
	validateAddrArray(waitValidateAddrArray)
}

func getValidateAddrArray(tcpAddrChan chan []*net.TCPAddr) []*net.TCPAddr {
	consumed := 0
	waitValidateAddrMap := make(map[int64]*net.TCPAddr, 0)
	for addrArray := range tcpAddrChan {
		consumed++
		for _, addr := range addrArray {
			if addr == nil || addr.IP == nil {
				continue
			}
			//FIXME : 现在只考虑了IPV4地址
			if addr.IP.DefaultMask() != nil {
				waitValidateAddrMap[utils.GetAddrHash(addr)] = addr
			}
		}
		if consumed == len(DNS_SEEDS) {
			break
		}
	}
	waitValidateAddrArray := make([]*net.TCPAddr, 0)
	for _, value := range waitValidateAddrMap {
		waitValidateAddrArray = append(waitValidateAddrArray, value)
	}
	return waitValidateAddrArray
}
func validateAddrArray(addrArray []*net.TCPAddr) {

	for _, addr := range addrArray {
		if addr == nil || addr.IP == nil {
			continue
		}
		go ValidateAddr(addr, func(validateAddr *net.TCPAddr, flag bool) {
			if flag {
				utils.GetSugarLogger().Infof("Get useful node %s", validateAddr.IP.String())
				if UsefulNodesHolder == nil {
					UsefulNodesHolder = &UsefulNodes{}
				}
				if !UsefulNodesHolder.Contain(validateAddr) {
					UsefulNodesHolder.Append(validateAddr)
				}
			} else {
				UsefulNodesHolder.FoundAndRemove(validateAddr)
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
		callback(addr, false)
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
		callback(addr, false)
		return
	}
	utils.GetSugarLogger().Infof("connect to %s", addr.IP.String())
	callback(addr, true)
}

