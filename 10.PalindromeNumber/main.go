package isMatch

import "fmt"

func isMatch(s string, p string) bool {
	sSize := len(s)
	pSize := len(p)
	fmt.Println(sSize, pSize)
	return len(s) == 0
}
