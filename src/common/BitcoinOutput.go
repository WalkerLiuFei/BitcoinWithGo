package common

import (
	"bytes"
	"encoding/binary"
	"io"
)

type BitcoinOuput struct {
	//input stream
	Stream []byte
}

func (output *BitcoinOuput) WriteNum(num interface{}) {
	var writer io.Writer = new(bytes.Buffer)
	binary.Write(writer, binary.LittleEndian, num)
	writer.Write(output.Stream)
}

func (output *BitcoinOuput) WriteString(str string) *BitcoinOuput {
	strBytes := []byte(str)
	output.WriteVarInt(int64(len(strBytes)))
	byteArray := []byte(str)
	output.Stream = append(output.Stream, byteArray...)
	return output
}
func (output *BitcoinOuput) WriteVarInt(num int64) *BitcoinOuput {
	if num < 0XFD {
		output.Stream = make([]byte, 1)
		output.Stream[0] = byte(num)
	} else if num <= 0xFFFF {
		output.Stream = make([]byte, 3)
		output.Stream[0] = 0XFD
		var writer io.ReadWriter = new(bytes.Buffer)
		binary.Write(writer, binary.LittleEndian, int16(num))
		writer.Read(output.Stream[1:])
	} else if num <= 0xffffffff {
		output.Stream = make([]byte, 5)
		output.Stream[0] = 0XFE
		var writer io.ReadWriter = new(bytes.Buffer)
		binary.Write(writer, binary.LittleEndian, int32(num))
		writer.Read(output.Stream[1:])
	} else {
		output.Stream = make([]byte, 9)
		output.Stream[0] = 0XFE
		var writer io.ReadWriter = new(bytes.Buffer)
		binary.Write(writer, binary.LittleEndian, num)
		writer.Read(output.Stream[1:])
	}
	return output
}
