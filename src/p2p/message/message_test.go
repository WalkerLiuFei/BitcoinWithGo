package message

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"utils"
	"strings"
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

func TestHeaderHashConsensus(t *testing.T) {
	msgHeader := "02000000" +
		"b6ff0b1b1680a2862a30ca44d346d9e8" +
		"910d334beb48ca0c0000000000000000" +
		"9d10aa52ee949386ca9385695f04ede2" +
		"70dda20810decd12bc9b048aaab31471" +
		"24d95a54" +
		"30c31b18" +
		"fe9f0864"
	//msgHeader := "04000000b9e2784a84e5d2468cee60ad14e08d0fee5dda49a37148040000000000000000e9dd2b13157508891880ef68729a1e5ecdde58062ebfa214a89f0141e5a4717faefd2b577627061880564bec"
	header := generateHexArray(msgHeader)
	fmt.Println(utils.ValidateHeaderHash(header))
}

func TestDecodeMessage(t *testing.T) {

}

func TestStringTrim(t *testing.T) {

	str := "******"
	fmt.Println(strings.TrimRightFunc(str, func(r rune) bool {
		return !(r > 96 && r < 123)
	}))
}