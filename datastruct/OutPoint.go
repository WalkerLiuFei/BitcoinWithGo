package datastruct

type OutPoint struct {
	/**
		32 字节,对应交易信息的hash值
	 */
	hash []byte

	/**
		32字节的无符号数，这笔交易在输出的位置，默认从零开始
	 */
	index int64
}
