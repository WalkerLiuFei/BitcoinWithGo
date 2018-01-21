package datastruct

import "common"

type OutPoint struct {
	/**
	32 字节,对应交易信息的hash值
	*/
	Hash []byte

	/*
	 *	32字节的无符号数，这笔交易在 输入 的位置，默认从零开始
	 */
	Index uint32
}

func (outPoint *OutPoint) Init(input common.BitcoinInput) {
	point := &OutPoint{}
	point.Hash = make([]byte, 32)
	input.ReadBytes(point.Hash)
	input.ReadNum(point.Index)
}

func (outPoint *OutPoint) GetBytes(input common.BitcoinInput) []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(outPoint.Hash).WriteNum(outPoint.Index)
	return output.Buffer.Bytes()
}
