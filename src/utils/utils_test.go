package utils

import (
	"testing"
	"fmt"
)

func TestGetLocalServiceAddr(t *testing.T) {
	fmt.Println(GetLocalServiceAddr())
}

func TestGetTargetThreshold(t *testing.T) {
	r := GetPreciseTargetByNum(0X181BC330)
	fmt.Println(len(r), r)
}
func TestShift(t *testing.T) {
	bits := []byte{0x18, 0x1b, 0xc3, 0x30}
	fmt.Printf("%x", int64(bits[1])<<uint32(16))
}
func TestGetPreciseTarget(t *testing.T) {
	bits := []byte{0x18, 0x1b, 0xc3, 0x30}
	r, err := GetPreciseTarget(bits)
	checkerr(err)
	fmt.Println(len(r), r)
}
func TestGetBytesHexString(t *testing.T) {
	byteArr := []byte{0x10, 0xBE, 0X52}
	fmt.Println(GetBytesHexString(byteArr))
}