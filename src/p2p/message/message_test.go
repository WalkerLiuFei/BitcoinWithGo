package message

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func generateHexArray(str string) []byte {
	input := []byte(str)
	if len(input)%2 != 0 {
		println("字符不是成对出现的,请检查")
	}
	result := make([]byte, 0)

	for index := 0; index < len(input); index += 2 {
		i64, _ := strconv.ParseInt(string(input[index:index+2]), 16, 0)
		result = append(result, byte(i64))
	}
	return result
}
func TestReject_Message_Decode(t *testing.T) {
	msgStr := "02747812156261642d74786e732d696e707574732d7370656e7439471" +
		"5fcab51093be7bfca5a31005972947baf86a31017939575fb2354222821"
	input := generateHexArray(msgStr)
	rejectMsg := Reject_Message{}
	rejectMsg.Decode(input)
	jsonStr, _ := json.Marshal(rejectMsg)
	fmt.Println(string(jsonStr))
}

func TestVersion_Message_Decode(t *testing.T) {
	msgStr := "721101000100000000000000bc8f5e5400000000010000000000000000000000000000000000" +
		"ffffc61b6409208d010000000000000000000000000000000000ffffcb0071c0208d128035cbc97953" +
		"f80f2f5361746f7368693a302e392e332fcf05050001"
	input := generateHexArray(msgStr)
	versionMsg := version_Message{}
	versionMsg.Decode(input)
	jsonStr, _ := json.Marshal(versionMsg)
	fmt.Println(string(jsonStr))
}

func TestGetBlockMessage_Decode(t *testing.T) {
	msgStr := "7111010002d39f608a7775b537729884d4e6633bb2" +
		"105e55a16a14d31b0000000000000000" +
		"5c3e6403d40837110a2e8afb602b1c01" +
		"714bda7ce23bea0a0000000000000000" +
		"00000000000000000000000000000000" +
		"00000000000000000000000000000000"
	input := generateHexArray(msgStr)
	getBlockMsg := GetBlockMessage{}
	getBlockMsg.Decode(input)
	jsonStr, _ := json.Marshal(getBlockMsg)
	fmt.Println(string(jsonStr))
}
