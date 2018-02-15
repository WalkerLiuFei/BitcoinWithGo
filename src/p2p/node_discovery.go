package p2p

import (
	"configs"
	"github.com/spf13/viper"
	"net"
	"time"
	"io/ioutil"
	"utils"
	"p2p/message"
	"github.com/syndtr/goleveldb/leveldb"
)

var DNS_SEEDS = []string{"bitseed.xf2.org", "dnsseed.bluematt.me", "seed.bitcoin.sipa.be", "dnsseed.bitcoin.dashjr.org", "seed.bitcoinstats.com"}

type validateCallback func(bool)

//更新本地有效的Node地址
func UpdateUsefulNode(peerChannel chan *net.TCPAddr) {
	//在直接通过DNS Seed之前应该先查询持久化的Node节点来通信
	tcpAddrChan := make(chan []*net.TCPAddr)
	ConsumeDNSSeed(tcpAddrChan)
	select {
	case addrArray := <-tcpAddrChan:
		db, _ := leveldb.OpenFile(viper.GetString(configs.DB_PATH), nil)
		defer db.Close()
		for _, addr := range addrArray {
			if addr == nil || addr.IP == nil {
				continue
			}

			go validateAddr(addr, func(flag bool) {
				if flag {

				}
			})
		}
	}

}

func ConsumeDNSSeed(chan [](*net.TCPAddr)) {
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
	}
	//return result
}

//链接成功返回true
func validateAddr(addr *net.TCPAddr, callback validateCallback) (*PeerConn) {
	conn, err := net.DialTCP("tcp", nil, addr)
	if conn == nil {
		callback(nil, false)
	}
	//defer Conn.Close()
	//FIXME : 连接是不是必须要关掉？
	checkerr(err, validateAddr)
	verMsg := message.InitVersionMessage(viper.GetInt(configs.CURRENT_HEIGHT), addr)
	var versionAckMsg []byte
	//等待11S
	for count := 0; count < 11 && len(versionAckMsg) == 0; count++ {
		_, err = conn.Write(verMsg.Encode())
		checkerr(err, validateAddr)
		utils.GetSugarLogger().Infof("connect to %s", addr.IP.String())
		time.Sleep(time.Second * 1)
		versionAckMsg, err = ioutil.ReadAll(conn)
	}
	checkerr(err, validateAddr)
	msg, err := message.DecodeMessage(versionAckMsg)
	checkerr(err, validateAddr)
	if msg == nil {
		return nil, false
	}
	//FIXME : 在做成连接池后,优化打印的信息

	return &PeerConn{
		Addr: addr,
		Conn: conn,
	}, true
}
