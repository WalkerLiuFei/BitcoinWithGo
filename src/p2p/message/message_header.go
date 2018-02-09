package message

import (
	"common"
	"go.uber.org/zap"
	"utils"
	"constants"
)


// reference : https://bitcoin.org/en/developer-reference#message-headers
type message_header struct {
	//魔法字节，标识
	MAGIC uint32
	//标识信息的类型,size 固定为12
	Command []byte
	//Payload size
	PayloadSize uint32
	//First 4 bytes of SHA256(SHA256(payload)) in
	CheckSum []byte
}

var logger, _ = zap.NewProduction()

func (msg *message_header) init(cmd constants.MessageType, payload []byte) {
	msg.MAGIC = 0xd9b4bef9
	msg.Command = []byte(cmd)
	if len(msg.Command) < 12 {
		index := len(msg.Command)
		for index < 12 {
			msg.Command = append(msg.Command, 0)
			index++
		}
	}
	msg.PayloadSize = uint32(len(payload))
	msg.CheckSum = utils.DoubleHash(payload)[:4]
}

func (msg *message_header) getBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(msg.MAGIC).
		WriteBytes(msg.Command).
		WriteNum(msg.PayloadSize).
		WriteBytes(msg.CheckSum)
	return output.Buffer.Bytes()
}

func checkError(e error) {
	if e != nil {
		utils.GetSugarLogger().Error(e.Error())
	}
}