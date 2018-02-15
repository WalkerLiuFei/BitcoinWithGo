package utils

import (
	"bytes"
	"fmt"
	"strconv"
)

func GenerateHexBytes(str string) []byte {
	input := []byte(str)
	if len(input)%2 != 0 {
		println("字符不是成对出现的,请检查")
	}
	result := make([]byte, 0)

	for index := 0; index < len(input); index += 2 {
		i64, _ := strconv.ParseInt(string(input[index:index+2]), 16, 0)
		result = append(result, byte(i64))
	}
	return result
}
//
func GetBytesHexString(src []byte) string {
	if len(src) == 0 {
		return ""
	}
	buffer := new(bytes.Buffer)
	for _, b := range src {
		fmt.Fprintf(buffer, "%x", b)
	}
	return string(buffer.Bytes())
}

func ReverseArray(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	for index, _ := range src {
		src[index], src[len(src)-index-1] = src[len(src)-index-1], src[index]
	}
	return src
}
