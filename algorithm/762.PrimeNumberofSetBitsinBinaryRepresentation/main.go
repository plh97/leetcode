package countPrimeSetBits

import "fmt"

func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		fmt.Println(i)
		cp := i
		for cp > 0 {
			if cp%2 == 1 {
				res++
			}
			cp /= 2
		}
	}
	return res
}
