package rotateString

import (
	"strings"
)

func rotateString(A string, B string) bool {
	if len(A) == len(B) && strings.Index(A+A, B) > -1 {
		return true
	}
	return false
}
