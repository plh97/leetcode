package fairCandySwap

func fairCandySwap(A []int, B []int) []int {
	res := []int{}
	durNum := total(A) - total(B)
	for i := range A {
		for j := range B {
			if (A[i]-B[j])*2 == durNum {
				return []int{A[i], B[j]}
			}
		}
	}
	return res
}

func total(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}
