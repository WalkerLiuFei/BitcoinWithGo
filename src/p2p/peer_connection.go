package p2p

import (
	"net"
	"p2p/message"
	"github.com/spf13/viper"
	"configs"
	"io/ioutil"
	"fmt"
	"utils"
	"time"
)

type PeerConn struct {
	//Peer 的地址
	IP *net.TCPAddr

	//
}
//链接成功返回true
func ConnectPeer(addr *net.TCPAddr) (*net.TCPConn, bool) {
	conn, err := net.DialTCP("tcp4", nil, addr)
	if conn == nil {
		return conn, false
	}
	//FIXME : 连接是不是必须要关掉？
	checkerr(err, ConnectPeer)
	verMsg := message.InitVersionMessage(viper.GetInt(configs.CURRENT_HEIGHT), addr)
	_, err = conn.Write(verMsg.Encode())
	checkerr(err, ConnectPeer)
	utils.GetSugarLogger().Infof("connect to %s", addr.IP.String())
	var result []byte
	for ; len(result) == 0; {
		//等待100ms
		time.Sleep(time.Millisecond * 100)
		result, err = ioutil.ReadAll(conn)
		checkerr(err, ConnectPeer)
		//FIXME : 在做成连接池后,优化打印的信息
	}
	fmt.Println(result)
	return conn, true
}
func checkerr(e error, funcName interface{}) {
	logger := utils.GetSugarLogger()
	if e != nil {
		logger.Errorf("call func name is %s,msg is %s", utils.GetFunctionName(funcName), e.Error())
	}

}
