package p2p

import (
)

//同步本地的区块链
func SyncBlockChain() {
	//FIXME:是不是要移到main方法里面执行？
	go UpdateUsefulNode()
	//addr := GetRandomNodeAsync()

}

/*
func GetHeaderMessage(conn *net.TCPAddr) []message.HeaderMessage {
	if conn == nil {
		return nil
	}
	//FIXME : 连接是否要在这里关闭？
	getHeaderMsg := &message.GetHeaderMessage{}
	getHeaderMsg.Init(nil, utils.GenerateHexBytes(viper.GetString(configs.CURRENT_TOP_BLOCK_HASH)))
	//conn.Write(getHeaderMsg.Encode())
	var result []message.HeaderMessage
	for count := 0; count < 500; count++ {
		//等待100ms
		time.Sleep(time.Millisecond * 100)
		//response, err := ioutil.ReadAll(conn)
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
*/