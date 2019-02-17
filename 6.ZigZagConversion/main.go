package main

import "fmt"

// 有意思啊.传值还是传指针, very fun
func coverStringIntoGraph(s string, r int) *[][]int {
	fmt.Println(s)
	graph:=&[][]int{}

	return graph
}

func convert(s string, numRows /*总列数*/ int) string {
	container := [][]string{}
	numCols := len(s) / numRows // 总行数
	numMid := numRows - 2
	graph := coverStringIntoGraph(s, numRows)
	fmt.Println(s, numCols, container, numMid,graph)
	return s
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
