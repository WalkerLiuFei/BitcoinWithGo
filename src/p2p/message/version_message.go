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
	header messageHeader
	//The highest protocol version understood by the transmitting node.
	version int32

	//The services supported by the transmitting node encoded as a bitfield.
	services constants.ServiceType

	//The current Unix epoch time according to the transmitting node’s clock.
	timestamp int64

	recipientAddress *datastruct.NetworkAddress

	senderAddress *datastruct.NetworkAddress

	//A random nonce which can help a node detect a connection to itself. If the nonce is 0, the nonce field is ignored. If the nonce is anything else, a node should terminate the connection on receipt of a version message with a nonce it previously sent.
	nonce uint64

	//required if user agent byte > 0,
	userAgent string

	startHeight int32

	//Transaction relay flag. If 0x00, no inv messages or tx messages announcing new transactions should be sent to this client until it sends a filterload message or filterclear message. If the relay field is not present or is set to 0x01, this node wants inv messages and tx messages announcing new transactions.
	relay bool
}

func InitVersionMessage(startHeight int, recipientAddr *net.TCPAddr) *version_Message {
	versionMsg := &version_Message{}
	versionMsg.version = constants.PROTOCOL_VERSION
	versionMsg.services = constants.NODE_NETWORK
	versionMsg.timestamp = time.Now().Unix()
	versionMsg.startHeight = int32(startHeight)
	versionMsg.recipientAddress = datastruct.InitByTCPAddr(recipientAddr, constants.NODE_NETWORK)
	versionMsg.senderAddress = datastruct.InitByTCPAddr(utils.GetLocalServiceAddr(), constants.NODE_NETWORK)
	versionMsg.nonce = uint64(viper.GetInt64(configs.NODE_ID))
	versionMsg.userAgent = constants.USER_AGENT
	versionMsg.relay = true
	//在初始化完Payload后再根据Payload来初始化Header
	versionMsg.header.init(constants.VERSION, versionMsg.GetPayload())
	return versionMsg
}

func (versionMsg *version_Message) Decode(contentBytes []byte) {
	input := common.NewBitcoinInput(contentBytes)
	versionMsg.header.decode(input)
	input.ReadNum(&versionMsg.version)
	input.ReadNum(&versionMsg.services)
	input.ReadNum(&versionMsg.timestamp)
	versionMsg.recipientAddress = datastruct.InitNetWorkAddress(input)
	versionMsg.senderAddress = datastruct.InitNetWorkAddress(input)
	input.ReadNum(&versionMsg.nonce)
	userAgentBytesCount, _ := input.ReadVarInt()
	userAgentBytes := make([]byte, userAgentBytesCount)
	input.ReadBytes(userAgentBytes)
	versionMsg.userAgent = string(userAgentBytes)
	input.ReadNum(&versionMsg.startHeight)
	input.ReadNum(&versionMsg.relay)
}

func (versionMsg *version_Message) Encode() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(versionMsg.header.getBytes()).WriteBytes(versionMsg.GetPayload())
	return output.Buffer.Bytes()
}
func (versionMsg *version_Message) GetPayload() []byte {
	output := common.BitcoinOuput{}
	output.WriteNum(versionMsg.version).
		WriteNum(versionMsg.services).
		WriteNum(versionMsg.timestamp).
		WriteBytes(versionMsg.recipientAddress.Encode()).
		WriteBytes(versionMsg.senderAddress.Encode()).
		WriteNum(versionMsg.nonce).
		WriteVarInt(int64(len(versionMsg.userAgent))).
		WriteBytes([]byte(versionMsg.userAgent)).
		WriteNum(versionMsg.startHeight).
		WriteNum(versionMsg.relay)
	return output.Buffer.Bytes()
}
