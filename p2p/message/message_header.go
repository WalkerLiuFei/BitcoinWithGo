package message

import (
	"common"
	"constants"
	"go.uber.org/zap"
	"utils"
	"github.com/spf13/viper"
	"configs"
)

// reference : https://bitcoin.org/en/developer-reference#message-headers
type messageHeader struct {
	//魔法字节，标识
	magic uint32
	//标识信息的类型,size 固定为12
	command []byte
	//Payload size
	payloadSize uint32
	//First 4 bytes of SHA256(SHA256(payload)) in
	checkSum []byte
}

var logger, _ = zap.NewProduction()

func (msg *messageHeader) init(cmd constants.MessageType, payload []byte) {
	msg.magic = uint32(viper.GetInt(configs.MAGIC))
	msg.command = []byte(cmd)
	if len(msg.command) < 12 {
		index := len(msg.command)
		for index < 12 {
			msg.command = append(msg.command, 0)
			index++
		}
	}
	msg.payloadSize = uint32(len(payload))
	msg.checkSum = utils.DoubleHash(payload)[:4]
}
func (msg *messageHeader) decode(input common.BitcoinInput) {
	input.ReadNum(&msg.magic)
	input.ReadBytes(msg.command)
	input.ReadNum(&msg.payloadSize)
	input.ReadBytes(msg.checkSum)
}

func (msg *messageHeader) getBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(msg.magic).
		WriteBytes(msg.command).
		WriteNum(msg.payloadSize).
		WriteBytes(msg.checkSum)
	return output.Buffer.Bytes()
}

func decodeHeader(inputStream []byte) (*messageHeader) {
	input := common.NewBitcoinInput(inputStream)
	//读取Message Header
	header := &messageHeader{}
	input.ReadNum(&header.magic)
	header.command = make([]byte, 12)
	input.ReadBytes(header.command)
	input.ReadNum(&header.payloadSize)
	header.checkSum = make([]byte, 4)
	input.ReadNum(&header.checkSum)
	return header
}

func checkError(e error) {
	if e != nil {
		utils.GetSugarLogger().Error(e.Error())
	}
}
