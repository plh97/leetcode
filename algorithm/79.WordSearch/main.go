package exist

import "strings"

func exist(board [][]byte, word string) bool {
	c := ""
	for _, col := range board {
		for _, e := range col {
			c += string(e)
		}
	}
	for _, e := range word {
		ii := strings.Index(c, string(e))
		if ii == -1 {
			return false
		} else {
			c = c[:ii] + c[ii+1:]
		}
	}
	return true
}
