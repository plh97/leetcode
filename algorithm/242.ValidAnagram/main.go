package isAnagram

import (
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		indexT := strings.Index(t, s[i:i+1])
		if indexT > -1 {
			t = t[:indexT] + t[indexT+1:]
		}
	}
	return len(t) == 0
}
