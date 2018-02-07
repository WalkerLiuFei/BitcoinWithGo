package configs

import (
	"github.com/spf13/viper"
	"math/rand"
	"constants"
)

const (
	//运行环境： test，localtest,realse
	CONTEXT_NETWORK_PORT = "context"

	//节点类型
	NODE_TYPE = "node-type"

	//node id
	NODE_ID = "node-id"

	//当前区块链高度
	CURRENT_HEIGHT = "current-height"
)

func InitConfigs() {
	//运行环境
	viper.Set(CONTEXT_NETWORK_PORT, constants.MAIN_NET_PORT)
	//节点类型
	viper.Set(NODE_TYPE, constants.NODE_NETWORK)
	//节点的ID ： 一个64位随机数
	viper.Set(NODE_ID, rand.Uint64())
	//TODO:当前节点的区块链高度,需要持久化的配合
	viper.Set(CURRENT_HEIGHT, 0)
}
