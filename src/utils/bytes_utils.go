package utils

import (
	"bytes"
	"fmt"
)

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
