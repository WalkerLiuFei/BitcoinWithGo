package configs

import (
	"constants"
	"github.com/spf13/viper"
	"math/rand"
	"runtime"
)

const (

	//运行环境： test，localtest,realse
	CONTEXT = "context"
	//端口号
	PORT = "port"
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

	//
	CURRENT_TOP_BLOCK_HASH = "current-top-block-hash"
	//数据库相对地址
	DB_PATH = "/"
)

func InitConfigs() {
	//
	viper.Set(CONTEXT, constants.MAIN_NET.Network)
	//运行环境
	viper.Set(PORT, uint16(constants.MAIN_NET.DefaultPort))
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
	//TODO : 需要持久化的配合记录
	viper.Set(CURRENT_TOP_BLOCK_HASH, constants.GENESIS_BLOCK_HASH)
	//FIXME :这个数据库文件保存的地址应该是由用户来定义的，目前先写死
	if runtime.GOOS == "windows" {
		viper.Set(DB_PATH, "g:/temp")
	}
	//TODO ：Linux path to store
}
