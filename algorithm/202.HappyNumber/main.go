package isHappy

func isHappy(n int) bool {
	table := []int{}
	for true {
		temp := 0
		for n > 0 {
			temp += (n % 10) * (n % 10)
			n /= 10
		}
		if indexOf(table,temp) > -1 {
			return false
		}
		table = append(table, temp)
		n = temp
		if temp == 1 {
			return true
		}
	}
	return false
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}



// solution 2
// func isHappy(n int) bool {
// 	for i := 0; i < 7; i++ {
// 		temp := 0
// 		for n > 0 {
// 			temp += (n % 10) * (n % 10)
// 			n /= 10
// 		}
// 		n = temp
// 		if temp == 1 {
// 			return true
// 		}
// 	}
// 	return false
// }