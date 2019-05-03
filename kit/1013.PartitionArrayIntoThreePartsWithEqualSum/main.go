package canThreePartsEqualSum

func canThreePartsEqualSum(A []int) bool {
	average := 0
	for i := range A {
		average += A[i]
	}
	average = average / 3
	l := 1
	for ; l < len(A)-1; l++ {
		a := A[:l]
		if countFunc(a) == average {
			break
		}
		if l == len(A)-2 {
			return false
		}
	}

	r := len(A) - 1

	for ; r > 0; r-- {
		a := A[r:]
		if countFunc(a) == average {
			break
		}
		if r == 1 {
			return false
		}
	}
	return l < r
}

func countFunc(nums []int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}
