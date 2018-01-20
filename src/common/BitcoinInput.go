package common

import (
	"io"
	"fmt"
	"go.uber.org/zap"
	"errors"
	"encoding/binary"
	"bytes"
)

var logger, _ = zap.NewProduction()

type BitcoinInput struct {
	//input stream
	Stream []byte
}

//type BitcoinInput interface{}

func (input *BitcoinInput) New(stream []byte) *BitcoinInput {
	input.Stream = stream
	return input
}

func (input *BitcoinInput) NewWithBuffer(buffer io.Reader) *BitcoinInput {

	_, err := buffer.Read(input.Stream)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println()
	return input
}

func (input *BitcoinInput) ReadString() (string) {
	streamLen, _ := input.ReadVarInt()
	if streamLen == 0 {
		return ""
	}
	bytes := input.Stream
	if int64(len(input.Stream)) > streamLen {
		bytes = input.Stream[:streamLen]
	}
	return string(bytes)
}

func (input *BitcoinInput) ReadVarInt() (int64, error) {
	if len(input.Stream) == 0 {
		return 0, errors.New("Cross EOF")
	}
	byte1 := 0XFF & input.Stream[0]
	if byte1 < 0XFD {
		return int64(byte1), nil
	}
	if byte1 == 0XFD {
		var num int16
		err := input.ReadNum(num)
		return int64(num), err
	}
	if byte1 == 0XFE {
		var num int32
		err := input.ReadNum(num)
		return int64(num), err
	}
	var num int64
	err := input.ReadNum(num)
	return num, err
}

func (input *BitcoinInput) ReadNum(numPointer interface{}) error {
	err := binary.Read(bytes.NewBuffer(input.Stream), binary.LittleEndian, numPointer)

	return err
}
