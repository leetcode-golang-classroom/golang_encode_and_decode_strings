package sol

import (
	"fmt"
	"strconv"
)

type ListHandle struct {
	sep string
}

func Constructor(sep string) ListHandle {
	return ListHandle{
		sep: sep,
	}
}
func (h *ListHandle) encode(list []string) string {
	encodeString := ""
	for _, val := range list {
		encodeString += fmt.Sprintf("%d%s%s", len(val), h.sep, val)
	}
	return encodeString
}

func (h *ListHandle) decode(str string) []string {
	decodeResult := []string{}
	pos := 0
	strLen := len(str)
	tempLen := ""
	for pos < strLen {
		if string(str[pos]) != h.sep {
			tempLen += string(str[pos])
		} else { // first meet #
			num, _ := strconv.Atoi(tempLen)
			decodeResult = append(decodeResult, str[pos+1:pos+num+1])
			tempLen = ""
			pos += num + 1
			continue
		}
		pos++
	}
	return decodeResult
}
