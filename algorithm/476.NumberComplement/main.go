package findComplement

func findComplement(num int) int {
	temp := num
	temp |= temp >> 1
	temp |= temp >> 2
	temp |= temp >> 4
	temp |= temp >> 8
	temp |= temp >> 16

	return num ^ temp
}

// func findComplement(num int) int {
// 	temp := 1
// 	for ; temp <= num; temp *= 2 {
// 	}
// 	return (temp - 1) ^ num
// }
