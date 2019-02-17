package main

import "fmt"

func convert(s string, numRows /*总列数*/int) string {
	//
	container := [][]string{}

	numCols := len(s) / numRows	// 总行数
	numMid := numRows - 2
	fmt.Println(s, numCols, container, numMid)
	return s
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
