package datastruct

import (
	"testing"
	//"io/ioutil"
	"io/ioutil"
	"common"
	"github.com/stretchr/testify/assert"
	"fmt"
	"utils"
)

func TestBlock_Input(t *testing.T) {
	_, err := ioutil.ReadFile("./block-000000000000000000f061205567dc79c4e718209a568879d66132e016968ac6.dat")
	if err != nil {
		logger.Error(err.Error())
	}
	input := common.BitcoinInput{}
	block := Block{}
	block.Init(input)
	assert.Equal(t, 351, len(block.TXns))
	fmt.Println(block.Header)
}
func TestBlock_Out(t *testing.T) {
	stream,err := ioutil.ReadFile("./block-000000000000000000f061205567dc79c4e718209a568879d66132e016968ac6.dat")
	if err != nil{
		logger.Error(err.Error())
	}
	input := common.NewBitcoinInput(stream)
	block := Block{}
	block.Init(input)
	//	assert.Equal(t,351, len(block.TXns))
	fmt.Printf("%x \n", block.Header.Bits)
	fmt.Println(utils.GetTargetThreshold(block.Header.Bits))
	fmt.Println(utils.DoubleHash(block.Header.GetBytes()))
}