package utils

import "crypto/sha256"

func SHA256(src []byte) []byte {
	h := sha256.New()
	h.Write(src)
	return h.Sum(nil)
}

func DoubleHash(src []byte) []byte {
	round1 := SHA256(src)
	return SHA256(round1)
}