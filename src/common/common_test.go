package common

import (
	"testing"
	"fmt"
)

func TestBitcoinInput_New(t *testing.T) {
	input := BitcoinInput{Stream: []byte{0x01, 0x00, 0x00, 0x00}}
	var num uint32
	input.ReadNum(&num)
	fmt.Println(num)
}
func TestBitcoinOutput_New(t *testing.T) {
	str := "walker liu 19693"
	bitcoinOuput := BitcoinOuput{}
	bitcoinOuput.WriteString(str)
	fmt.Println(bitcoinOuput.Stream)

}
