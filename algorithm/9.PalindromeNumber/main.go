package isPalindrome

import (
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for x:= range str {
		if str[x:x+1] != str[len(str)-x-1:len(str)-x] {
			return false
		}
	}
	return true
}
