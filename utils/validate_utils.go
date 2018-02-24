package utils

import (
	"bytes"
	"errors"
	"math/big"
)

//根据官方文档计算的，通过nbit来计算Mining的Threshold : 这个是一个 256进制的数！
func GetTargetThreshold(nBits uint32) []byte {
	result := big.NewInt(int64(nBits & 0X00FFFFFF))
	power := (nBits & 0XFF000000 >> 24) - 3
	power *= 2
	for power > 0 {
		result = result.Mul(result, big.NewInt(16))
		power--
	}
	return result.Bytes()
}

//TODO : handler the sign bit ,if thr sign bit is set
func GetPreciseTarget(nBits []byte) ([]byte, error) {
	if len(nBits) != 4 {
		return nil, errors.New("nBits length is not right")
	}
	//256进制的byte数组
	var base = (int64(nBits[1]) << 16) + (int64(nBits[2]) << 8) + int64(nBits[3])
	result := big.NewInt(base)
	power := nBits[0] * 2
	for power > 0 {
		result = result.Mul(result, big.NewInt(256))
		power--
	}
	byteLen := len(result.Bytes())
	resultArr := make([]byte, 64)
	for index, _ := range resultArr {
		if index >= 64-byteLen {
			resultArr[index] = result.Bytes()[index+byteLen-64]
		}
	}
	return resultArr, nil
}

//获得一个准确的16进制的，一种512位长的Target难度数值
func GetPreciseTargetByNum(nBits uint32) []byte {
	//256进制的byte数组
	result := big.NewInt(int64(nBits & 0X00FFFFFF))
	power := (nBits & 0XFF000000 >> 24) - 3
	power *= 2
	for power > 0 {
		result = result.Mul(result, big.NewInt(256))
		power--
	}
	byteLen := len(result.Bytes())
	resultArr := make([]byte, 64)
	for index, _ := range resultArr {
		if index >= 64-byteLen {
			resultArr[index] = result.Bytes()[index+byteLen-64]
		}
	}
	return resultArr
}

func ValidateHeaderHash(header []byte) bool {
	if len(header) != 80 {
		return false
	}
	//actually 72 - 76,一共四个字节
	nBits := header[72:76]
	target, err := GetPreciseTarget(nBits)
	checkerr(err)
	headerHash := ReverseArray(DoubleHash(header))
	return bytes.Compare(headerHash, target) >= 0
}
