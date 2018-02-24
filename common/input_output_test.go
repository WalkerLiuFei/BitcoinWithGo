package common

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBitcoinInput_ReadNum(t *testing.T) {
	num := uint32(1000000)
	input := BitcoinInput{}
	output := &BitcoinOuput{}
	output.WriteNum(num)
	input.InputBuffer = new(bytes.Buffer)
	input.InputBuffer.Write(output.Buffer.Bytes())
	var num2 uint32
	input.ReadNum(&num2)
	fmt.Println(num2)
}
