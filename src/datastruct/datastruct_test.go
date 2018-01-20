package datastruct

import (
	"testing"
	"encoding/binary"
	"go.uber.org/zap"
	"io"
	"fmt"
	"bytes"
)

func TestDatastruct(t *testing.T) {
	var num int32 = -100
	var writer io.ReadWriter = new(bytes.Buffer)
	byteArr := make([]byte, 0)
	binary.Write(writer, binary.LittleEndian, num)
	writer.Write(byteArr)
	fmt.Println(byteArr)
	var result int32
	binary.Read(writer, binary.LittleEndian, &result)
	logger, _ := zap.NewProduction()
	logger.Info("the num is :", zap.Int32("value", result), )
}
