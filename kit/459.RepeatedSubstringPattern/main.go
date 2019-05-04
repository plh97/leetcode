package repeatedSubstringPattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		if isLoop(s[:i], s) {
			return true
		}
	}
	return false
}

func isLoop(a, b string) bool {
	for len(b) > 0 {
		index := strings.Index(b, a)
		if index == 0 {
			b = b[index+len(a):]
		} else {
			return false
		}
	}
	return true
}
