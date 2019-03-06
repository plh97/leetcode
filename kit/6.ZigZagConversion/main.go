package convert

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if len(s) < 3 {
		return s
	}
	ss := bytes.Buffer{}
	for i := 0; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
	}
	for i := 0; i < numRows-2; i++ {
		for j := 1; j < len(s)-i; j += (2*numRows - 2) {
			ss.WriteByte(s[j+i])
			if j+(2*numRows-2)-i-2 < len(s) {
				ss.WriteByte(s[j+2*numRows-i-4])
			}
		}
	}
	for i := numRows - 1; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
	}
	return ss.String()
}
