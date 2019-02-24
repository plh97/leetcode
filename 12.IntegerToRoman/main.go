package intToRoman

import "fmt"

// import "fmt"
// func main() {
// 	fmt.Println(isMatch("ippi", "p*."))
// 	// fmt.Println(isMatch("mississippi", "mis*is*ip*."))
// }
type Roman struct {
	Symbol string
	Value  int
}

func intToRoman(num int) string {
	res := ""
	romanMap := [...]Roman{
		{Symbol: "M", Value: 1000},
		{Symbol: "D", Value: 500},
		{Symbol: "C", Value: 100},
		{Symbol: "L", Value: 50},
		{Symbol: "X", Value: 10},
		{Symbol: "V", Value: 5},
		{Symbol: "I", Value: 1},
	}
	for i := range romanMap {
		fmt.Println(romanMap[i])
	}
	return res
}
