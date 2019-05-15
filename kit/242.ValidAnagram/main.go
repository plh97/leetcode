package isAnagram

import "fmt"

func isAnagram(s string, t string) bool {
	for i := range t {
		fmt.Println(s[i] + t[i])
	}
	return true
}
