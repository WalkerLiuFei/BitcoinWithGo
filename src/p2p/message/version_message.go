package message

import (
	"common"
	"configs"
	"constants"
	"datastruct"
	"github.com/spf13/viper"
	"net"
	"time"
	"utils"
)

//The version message provides information about the transmitting node to the receiving node at the beginning of a connection.
//reference : https://bitcoin.org/en/developer-reference#Version
type version_Message struct {
	Header message_header
	//The highest protocol version understood by the transmitting node.
	Version int32

	//The services supported by the transmitting node encoded as a bitfield.
	Services constants.ServiceType

	//The current Unix epoch time according to the transmitting node’s clock.
	TimeStamp int64

	RecipientAddress *datastruct.NetworkAddress

	SenderAddress *datastruct.NetworkAddress

	//A random nonce which can help a node detect a connection to itself. If the nonce is 0, the nonce field is ignored. If the nonce is anything else, a node should terminate the connection on receipt of a version message with a nonce it previously sent.
	Nonce uint64

	//required if user agent byte > 0,
	UserAgent string

	StartHeight int32

	//Transaction relay flag. If 0x00, no inv messages or tx messages announcing new transactions should be sent to this client until it sends a filterload message or filterclear message. If the relay field is not present or is set to 0x01, this node wants inv messages and tx messages announcing new transactions.
	Relay bool
}

func InitVersionMessage(startHeight int, recipientAddr *net.TCPAddr) *version_Message {
	versionMsg := &version_Message{}
	versionMsg.Version = constants.PROTOCOL_VERSION
	versionMsg.Services = constants.NODE_NETWORK
	versionMsg.TimeStamp = time.Now().Unix()
	versionMsg.StartHeight = int32(startHeight)
	versionMsg.RecipientAddress = datastruct.InitByTCPAddr(recipientAddr, constants.NODE_NETWORK)
	versionMsg.SenderAddress = datastruct.InitByTCPAddr(utils.GetLocalServiceAddr(), constants.NODE_NETWORK)
	versionMsg.Nonce = uint64(viper.GetInt64(configs.NODE_ID))
	versionMsg.UserAgent = constants.USER_AGENT
	versionMsg.Relay = true
	//在初始化完Payload后再根据Payload来初始化Header
	versionMsg.Header.init(constants.VERSION, versionMsg.GetPayload())
	return versionMsg
}

func (versionMsg *version_Message) Decode(payload []byte) {
	versionMsg.Header.init(constants.VERSION, payload)
	input := common.NewBitcoinInput(payload)
	input.ReadNum(&versionMsg.Version)
	input.ReadNum(&versionMsg.Services)
	input.ReadNum(&versionMsg.TimeStamp)
	versionMsg.RecipientAddress.Init(input)
	versionMsg.SenderAddress.Init(input)
	input.ReadNum(&versionMsg.Nonce)
	userAgentBytesCount, _ := input.ReadVarInt()
	userAgentBytes := make([]byte, userAgentBytesCount)
	input.ReadBytes(userAgentBytes)
	versionMsg.UserAgent = string(userAgentBytes)
	input.ReadNum(&versionMsg.StartHeight)
	input.ReadNum(&versionMsg.Relay)
}

func (versionMsg *version_Message) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(versionMsg.Header.getBytes()).WriteBytes(versionMsg.GetPayload())
	return output.Buffer.Bytes()
}
func (versionMsg *version_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(versionMsg.Version).
		WriteNum(versionMsg.Services).
		WriteNum(versionMsg.TimeStamp).
		WriteBytes(versionMsg.RecipientAddress.Encode()).
		WriteBytes(versionMsg.SenderAddress.Encode()).
		WriteNum(versionMsg.Nonce).
		WriteVarInt(int64(len(versionMsg.UserAgent))).
		WriteBytes([]byte(versionMsg.UserAgent)).
		WriteNum(versionMsg.StartHeight).
		WriteNum(versionMsg.Relay)
	return output.Buffer.Bytes()
}
