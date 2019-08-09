package queryString

import (
	"fmt"
	"strings"
)

func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		if !strings.Contains(S, fmt.Sprintf("%b", i)) {
			return false
		}
	}
	return true
}
