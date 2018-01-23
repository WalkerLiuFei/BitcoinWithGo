package common

import (
	"testing"
	"fmt"
	"bytes"
)

func TestBitcoinInput_New(t *testing.T) {
	inputBytes := []byte{0XFF, 0xFF, 0xFF, 0xFF, 0x10, 0x11}
	input := BitcoinInput{}
	input.New(inputBytes)
	var num uint32
	input.ReadNum(&num)
	fmt.Println(num)
	var p [2]byte
	input.ReadBytes(p)
	fmt.Println(p)
	input.ReadBytes(p)
	fmt.Println(p)
}
func TestBitcoinOutput_New(t *testing.T) {
	str := "walker liu 19693"
	bitcoinOuput := BitcoinOuput{}
	bitcoinOuput.WriteString(str)
	fmt.Println(bitcoinOuput.Buffer.Bytes())

}
