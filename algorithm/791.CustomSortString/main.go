package customSortString

import (
	"strings"
)

func customSortString(S string, T string) string {
	res := ""
	for i := range S {
		ii := -2
		for ii == -2 || ii > -1 {
			ii = strings.Index(T, S[i:i+1])
			if ii > -1 {
				res += S[i : i+1]
				T = T[:ii] + T[ii+1:]
			}
		}
	}
	return res + T
}

