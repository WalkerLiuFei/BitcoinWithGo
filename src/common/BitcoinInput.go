package common

import (
	"io"
	"go.uber.org/zap"
	"errors"
	"encoding/binary"
	"bytes"
)

var logger, _ = zap.NewProduction()

type BitcoinInput struct {
	InputBuffer *bytes.Buffer
}

//type BitcoinInput interface{}

func NewBitcoinInput(stream []byte) *BitcoinInput {
	input := &BitcoinInput{}
	input.InputBuffer = bytes.NewBuffer(stream)
	return input
}


func (input *BitcoinInput) NewWithBuffer(buffer io.Reader) *BitcoinInput {
	//input.InputBuffer = InputBuffer
	return input
}

func (input *BitcoinInput) ReadBytes(p []byte) {
	input.InputBuffer.Read(p)
}

func (input *BitcoinInput) ReadString() (string) {
	streamLen, _ := input.ReadVarInt()
	if streamLen == 0 {
		return ""
	}
	var p = make([]byte, streamLen)
	input.InputBuffer.Read(p)
	return string(p)
}

func (input *BitcoinInput) ReadVarInt() (int64, error) {
	if input.InputBuffer.Len() == 0 {
		return 0, errors.New("Cross EOF")
	}
	byte1, _ := input.InputBuffer.ReadByte()
	byte1 = byte1 & 0XFF
	if byte1 < 0XFD {
		return int64(byte1), nil
	}
	if byte1 == 0XFD {
		var num int16
		err := input.ReadNum(&num)
		return int64(num), err
	}
	if byte1 == 0XFE {
		var num int32
		err := input.ReadNum(&num)
		return int64(num), err
	}
	var num int64
	err := input.ReadNum(&num)
	return num, err
}

func (input *BitcoinInput) ReadNum(numPointer interface{}) error {
	err := binary.Read(input.InputBuffer, binary.LittleEndian, numPointer)
	return err
}
