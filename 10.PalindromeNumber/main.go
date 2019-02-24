package main

import "fmt"

func main() {
	fmt.Println(isMatch("aaaaaa", "a*a"))
}

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	dp := make([][]bool, slen-1)
	

	return dp[slen][plen]
}
