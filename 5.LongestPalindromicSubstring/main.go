package main

import "fmt"

func longestPalindromicSubstring(s string) string {
	fmt.Printf(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			fmt.Println(i, j)
		}
	}
	return s
}

func main() {
	fmt.Println(longestPalindromicSubstring("babad"))
}
