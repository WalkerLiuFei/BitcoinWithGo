package p2p

import (
	"configs"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"p2p/message"
	"time"
	"utils"
)

type PeerConn struct {
	//Peer 的地址
	Addr *net.TCPAddr
	//有效的连接
	Conn *net.TCPConn
}

func GetHeaderMessage(conn *net.TCPConn) []message.HeaderMessage {
	if conn == nil {
		return nil
	}
	//FIXME : 连接是否要在这里关闭？
	getHeaderMsg := &message.GetHeaderMessage{}
	getHeaderMsg.Init(nil, utils.GenerateHexBytes(viper.GetString(configs.CURRENT_TOP_BLOCK_HASH)))
	conn.Write(getHeaderMsg.Encode())
	var result []message.HeaderMessage
	for count := 0; count < 500; count++ {
		//等待100ms
		time.Sleep(time.Millisecond * 100)
		response, err := ioutil.ReadAll(conn)
		checkerr(err, GetHeaderMessage)
		if len(response) != 0 {
			headerMsg := message.HeaderMessage{}
			headerMsg.Decode(response)
			utils.GetSugarLogger().Infof("Get header message : %s", utils.GetBytesHexString(headerMsg.GetPayload()))
			result = append(result, headerMsg)
		}
	}
	return result
}
func checkerr(e error, funcName interface{}) {
	logger := utils.GetSugarLogger()
	if e != nil {
		logger.Errorf("call func name is %s,msg is %s", utils.GetFunctionName(funcName), e.Error())
	}
}
