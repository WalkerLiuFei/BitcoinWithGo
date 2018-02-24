package p2p

import (
	"net"
	"utils"
)

type PeerConn struct {
	//Peer 的地址
	Addr *net.TCPAddr
	//有效的连接
	Conn *net.TCPConn
}

func checkerr(e error, funcName interface{}) {
	logger := utils.GetSugarLogger()
	if e != nil {
		logger.Errorf("call func name is %s,msg is %s", utils.GetFunctionName(funcName), e.Error())
	}
}
