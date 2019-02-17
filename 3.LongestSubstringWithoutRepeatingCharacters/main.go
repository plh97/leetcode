package main

import "fmt"

func isRepeat(s string) bool {
	m := make(map[string]string)
	for i := 0; i < len(s); i++ {
		m[s[i:i+1]] = s[i : i+1]
	}
	return len(m) != len(s)
}

func lengthOfLongestSubstring(s string) int {
	m := 0
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isRepeat(s[i:j]) {
				// 重复就直接退出
				j = len(s)
			} else {
				if m < len(s[i:j]) {
					m = len(s[i:j])
				}
			}
		}
	}
	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}
