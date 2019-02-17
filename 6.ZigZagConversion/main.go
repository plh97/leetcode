package main

import "fmt"

func coverStringIntoGraph(s *string , r int) [][]int{
	return [][]int{}
}


func convert(s string, numRows /*总列数*/int) string {
	//
	container := [][]string{}

	numCols := len(s) / numRows	// 总行数
	numMid := numRows - 2
	coverStringIntoGraph()
	fmt.Println(s, numCols, container, numMid)
	return s
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
