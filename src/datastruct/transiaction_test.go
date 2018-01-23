package datastruct

import (
	"testing"
	"io/ioutil"
	"common"
	"fmt"
)

func TestTransaction(t *testing.T) {
	txData,err := ioutil.ReadFile("./tx-582a10734982c74693eadc53b7b1bdbed0840aeec568b6f890e685f08cf79473.dat")
	if err != nil{
		logger.Error(err.Error())
	}

	trans := Transaction{}
	input := common.BitcoinInput{}
	input.New(txData)
	trans.Init(input)
	fmt.Println(trans)

}
