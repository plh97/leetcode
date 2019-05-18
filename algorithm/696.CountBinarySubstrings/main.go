package countBinarySubstrings

func countBinarySubstrings(s string) int {
	res := []string{}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 2; j < len(s); j += 2 {
			if isEqual(s[i:j]) &&
				indexOf(res, s[i:j]) == -1 {
				res = append(res, s[i:j])
			}
			// fmt.Println(s[i:j],res)
		}
	}
	return len(res)
}

func isEqual(s string) bool {
	res := 0
	for i := range s {
		if s[i] == '0' {
			res--
		} else {
			res++
		}
	}
	return res == 0
}

func indexOf(nums []string, val string) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
