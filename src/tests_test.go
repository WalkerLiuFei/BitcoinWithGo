package src

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
