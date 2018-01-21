package common

import (
	"bytes"
	"encoding/binary"
	"io"
)

type BitcoinOuput struct {
	//input stream
	Buffer bytes.Buffer
}

func (output *BitcoinOuput) WriteNum(num interface{}) *BitcoinOuput {
	var writer io.Writer = new(bytes.Buffer)
	binary.Write(writer, binary.LittleEndian, num)
	var p []byte
	writer.Write(p)
	output.Buffer.Write(p)
	return output
}

func (output *BitcoinOuput) WriteString(str string) *BitcoinOuput {
	output.Buffer.WriteString(str)
	return output
}
func (output *BitcoinOuput) WriteBytes(byteArr []byte) *BitcoinOuput {
	output.Buffer.Write(byteArr)
	return output
}
func (output *BitcoinOuput) WriteVarInt(num int64) *BitcoinOuput {
	if num < 0XFD {
		output.Buffer.WriteByte(byte(num))
	} else if num <= 0xFFFF {
		output.Buffer.WriteByte(0XFD)
		output.WriteNum(int16(num))
	} else if num <= 0xFFFFFFFF {
		output.Buffer.WriteByte(0XFE)
		output.WriteNum(int32(num))
	} else {
		output.Buffer.WriteByte(0XFF)
		output.WriteNum(num)
	}
	return output
}
