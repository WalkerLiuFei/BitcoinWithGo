package utils

import (
	"crypto/sha256"
)

func SHA256(src []byte) []byte {
	h := sha256.New()
	h.Write(src)
	return h.Sum(nil)
}

func DoubleHash(src []byte) []byte {
	round1 := SHA256(src)
	return SHA256(round1)
}

func MerkleHash(hashes [][]byte) [][]byte {
	count := len(hashes) / 2
	extra := len(hashes) % 2
	results := make([][]byte, count+extra)
	for index := 0; index < count; index++ {
		dst := append(hashes[2*index], hashes[2*index+1]...)
		results[index] = DoubleHash(dst)
	}
	if extra == 1 {
		length := len(hashes) - 1
		dst := append(hashes[length], hashes[length]...)
		results[count] = DoubleHash(dst)
	}
	return results
}
