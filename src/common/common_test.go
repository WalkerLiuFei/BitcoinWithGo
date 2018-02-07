package common

import (
	"testing"
	"fmt"
)

func TestBitcoinInput_New(t *testing.T) {
}
func TestBitcoinOutput_New(t *testing.T) {
	str := "walker liu 19693"
	bitcoinOuput := BitcoinOuput{}
	bitcoinOuput.WriteString(str)
	fmt.Println(bitcoinOuput.Buffer.Bytes())

}
