package convert

import (
	"bytes"
)

// 有意思啊.传值还是传指针, very fun

func convert(s string, numRows /*总列数*/ int) string {
	ss := bytes.Buffer{}
	// 第一行
	for i := 0; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
	}
	// 中间行
	for i := 0; i < numRows-1; i++ {
		ss.WriteByte(s[i*(2*numRows - 2)+1])

	}
	// 第 2<=>n-1 行
	for i := numRows - 1; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
	}
	return ss.String()
}
