package isPowerOfFour

// 1
// 100
// 10000
// 1000000
// 规律是二进制表示中, 有且仅有1个是1,且位于奇数
func isPowerOfFour(num int) bool {
	temp := parseTwo2Ten("1010101010101010101010101010101")
	if num != 0 &&
		num&(num-1) == 0 &&
		num&temp == num {
		return true
	}
	return false
}
func parseTwo2Ten(s string) int {
	res := uint(0)
	for _, e := range s {
		res = res << 1
		if e == '1' {
			res++
		}
	}
	return int(res)
}

// func isPowerOfFour(num int) bool {
//     if num < 1 {
//         return false
//     }
// 	for i := uint(1); i <= uint(num); i = i << 2 {
// 		if i == uint(num) {
// 			return true
// 		}
// 	}
// 	return false
// }
