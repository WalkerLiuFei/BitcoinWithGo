package configs

import (
	"constants"
	"github.com/spf13/viper"
	"math/rand"
)

const (

	//运行环境： test，localtest,realse
	CONTEXT = "context"
	//端口号
	PORT = "context"
	//magic
	MAGIC = "magic"
	//节点类型
	NODE_TYPE = "node-type"
	//Max nbits
	MAX_N_BITS = "max-n-bits"
	//node id
	NODE_ID = "node-id"

	//当前区块链高度
	CURRENT_HEIGHT = "current-height"
)

func InitConfigs() {
	//
	viper.Set(CONTEXT, constants.MAIN_NET.Network)
	//运行环境
	viper.Set(PORT, constants.MAIN_NET.DefaultPort)
	//消息流的开头的四个字节
	viper.Set(MAGIC, constants.MAIN_NET.Magic)
	//最大的nbits
	viper.Set(MAX_N_BITS, constants.MAIN_NET.MaxNBits)
	//节点类型
	viper.Set(NODE_TYPE, constants.NODE_NETWORK)
	//节点的ID ： 一个64位随机数
	viper.Set(NODE_ID, rand.Uint64())
	//TODO:当前节点的区块链高度,需要持久化的配合
	viper.Set(CURRENT_HEIGHT, 0)
}
