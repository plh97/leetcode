package plusOne

// [1,2,9]
// [1,3,0]
func plusOne(digits []int) []int {
	addLevel := false
	for i := len(digits) - 1; i >= 0; i-- {

		// 如果需要进位
		if addLevel {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				addLevel = false
			}
		}
		// 首位
		if i == len(digits)-1 {
			if digits[len(digits)-1] == 9 {
				digits[len(digits)-1] = 0
				addLevel = true
			} else {
				digits[len(digits)-1]++
			}
		}
		// 如果是最后一位 , 但是 addLevel == true
		if addLevel && i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
	}
	return digits
}
