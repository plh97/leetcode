package main

import (
	"bytes"
	"fmt"
)

// 有意思啊.传值还是传指针, very fun

func convert(s string, numRows /*总列数*/ int) string {
	ss := bytes.Buffer{}
	// index := 0
	// for i := 0; i < numRows; i++ {
	// 	// fmt.Printf("%c %d\n", ss, (2*numRows - 2))
	// 	if i == 0 {
	// 		for j := 0; j < len(s)/numRows; j++ {
	// 			ss.WriteByte(s[(2*numRows-2)*j])
	// 			index++
	// 		}

	// 		for j := 0; j < len(s)/numRows; j++ {
	// 			if(index==len(s)-1) {
	// 				return ss.String()
	// 			}
	// 			fmt.Println(index)
	// 			// ss.WriteByte(s[(2*numRows-2)*j + numRows-2])
	// 			index++
	// 		}
	// 	}
	// }
	// 第一行
	for i := 0; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
		// index++
	}
	// 中间行
	for i := 0; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
		// index++
	}
	// 第 2<=>n-1 行
	for i := numRows - 1; i < len(s); i += (2*numRows - 2) {
		ss.WriteByte(s[i])
		// index++
	}

	return ss.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
