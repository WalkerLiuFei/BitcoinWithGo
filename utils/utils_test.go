package utils

import (
	"fmt"
	"testing"
	"time"
	"math/rand"
	"math"
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

/**
	基准测试： 我用的电脑： I7-4990处理器，集成显卡。测试下每秒可以处理多少次的Double Hash
	由于区块头的大小是80字节，我就随机产生80个字节值然后进行Double Hash计算
	在一块Radeon HD5830 显卡上每秒可以计算  622 million 次 SHA-256的计算
  我这个电脑 大约每秒可以计算150万次的 80字节的Double Hash-256

*/
func TestDoubleHashPerSecond(t *testing.T) {
	fmt.Println(time.Now().Unix())
	src := make([]byte, 80)
	fullfill(src)
	for count := 1000000; count >= 0; count-- {

		DoubleHash(src)
	}
	fmt.Println(time.Now().Unix())
}
func fullfill(src []byte) {
	for index, _ := range src {
		src[index] = byte(rand.Int31n(math.MaxInt16))
	}
}