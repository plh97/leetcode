package hasGroupsSizeX

import (
	"sort"
)

func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)
	res := [][]int{}
	for i := 0; i < deck[len(deck)-1]+1; i++ {
		res = append(res, []int{})
	}
	for i := range deck {
		res[deck[i]] = append(res[deck[i]], deck[i])
	}
	for i := 0; i < len(res); i++ {
		if len(res[i]) == 0 {
			res = append(res[:i], res[i+1:]...)
		}
	}

	for i := 0; i < len(res); i++ {
		if len(res[i]) != 0 && gcd(len(res[0]), len(res[i])) < 2 {
			return false
		}
	}
	return true
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
