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
	}
	r := len(A) - 1
	for ; r > 0; r-- {
		a := A[r:]
		if countFunc(a) == average {
			break
		}
	}
	return l < r
}

func countFunc(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}
