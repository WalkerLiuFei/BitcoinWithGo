package datastruct

import (
	"testing"
	//"io/ioutil"
	"io/ioutil"
	"common"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestBlock_1(t  *testing.T)  {
	stream,err := ioutil.ReadFile("./block-000000000000000000f061205567dc79c4e718209a568879d66132e016968ac6.dat")
	if err != nil{
		logger.Error(err.Error())
	}
	input := common.BitcoinInput{}
	input.New(stream)
	block := Block{}
	block.Init(input)
	assert.Equal(t,351, len(block.TXns))
	fmt.Println(block.Header)
}