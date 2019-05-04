package firstUniqChar

import (
	"strings"
)

func firstUniqChar(s string) int {
	for i := range s {
		currentS := s[i : i+1]
		lastIndex := strings.LastIndex(s, currentS)
		firstIndex := strings.Index(s, currentS)
		if lastIndex == firstIndex {
			return i
		}
	}
	return -1
}
