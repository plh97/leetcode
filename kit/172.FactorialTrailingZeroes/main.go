package trailingZeroes

// 仅仅只需要计算
// Runtime: 2472 ms, faster than 5.48% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// func trailingZeroes(n int) int {
// 	if n < 5 {
// 		return 0
// 	}
// 	count := 0
// 	for base := 5; base <= n; base += 5 {
// 		if base%5 == 0 {
// 			count++
// 			bbase := base
// 			for bbase/5%5 == 0 {
// 				count++
// 				bbase /= 5
// 			}
// 		}
// 	}
// 	return count
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Factorial Trailing Zeroes.
func trailingZeroes(n int) int {
	c := 0
	for n > 0 {
		c += n / 5
		n /= 5
	}
	return c
}
