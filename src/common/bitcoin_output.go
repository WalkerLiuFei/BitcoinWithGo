package common

import (
	"bytes"
	"encoding/binary"
	"utils"
)

func (output *BitcoinOuput) WriteNum(num interface{}) *BitcoinOuput {
	if output.Buffer == nil {
		output.Buffer = new(bytes.Buffer)
	}
	err := binary.Write(output.Buffer, binary.LittleEndian, num)
	checkError(err)
	return output
}

type BitcoinOuput struct {
	//input stream
	Buffer *bytes.Buffer
}

func checkError(e error) {
	if e != nil {
		utils.GetSugarLogger().Error(e.Error())
	}

}

func (output *BitcoinOuput) WriteString(str string) *BitcoinOuput {
	if output.Buffer == nil {
		output.Buffer = new(bytes.Buffer)
	}
	output.Buffer.WriteString(str)
	return output
}
func (output *BitcoinOuput) WriteBytes(byteArr []byte) *BitcoinOuput {
	if output.Buffer == nil {
		output.Buffer = new(bytes.Buffer)
	}
	output.Buffer.Write(byteArr)
	return output
}
func (output *BitcoinOuput) WriteVarInt(num int64) *BitcoinOuput {
	if output.Buffer == nil {
		output.Buffer = new(bytes.Buffer)
	}
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
