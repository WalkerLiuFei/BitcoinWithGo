package persistence

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/spf13/viper"
	"configs"
	"utils"
	"net"
	"encoding/json"
)

const USEFUL_NODE = "useful-node"

func GetDBInstance() *leveldb.DB {
	//FIXME : level DB 有自定义的内存缓存，也就是内存缓冲，默认大小为8MB定义在第二个option选项中
	db, err := leveldb.OpenFile(viper.GetString(configs.DB_PATH), nil)
	checkError(err)
	return db
}
func checkError(err error) {
	if err != nil {
		utils.GetSugarLogger().Error(err.Error())
	}

}

func StoreUsefulNodes(addrArray []*net.TCPAddr) {
	utils.GetSugarLogger().Infof("Store the addr to db")
	db := GetDBInstance()
	defer db.Close()
	if db == nil {
		return
	}
	batch := new(leveldb.Batch)
	content, err := json.Marshal(addrArray)
	checkError(err)
	batch.Put([]byte(USEFUL_NODE), content)
	err = db.Write(batch, nil)
}
