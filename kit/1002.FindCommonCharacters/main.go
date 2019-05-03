package commonChars

import (
	"strings"
)

func commonChars(A []string) []string {
	res := []string{}
	for i := range A[0] {
		currentString := A[0][i : i+1]
		find := true
		for j := 1; j < len(A); j++ {
			if strings.Index(A[j], currentString) == -1 {
				find = false
				break
			}
		}
		if find {
			res = append(res, currentString)
			for j := 1; j < len(A); j++ {
				currentStringIndex := strings.Index(A[j], currentString)
				A[j] = A[j][:currentStringIndex] + A[j][currentStringIndex+1:]
			}
		}
	}
	return res
}
