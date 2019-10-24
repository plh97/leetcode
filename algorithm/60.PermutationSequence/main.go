package getpermutation

func getPermutation(n int, k int) string {
	res := ""
	length := 1
	kk := k
	nn := n
	arr := []int{}
	for kk > 0 {
		arr = append([]int{kk}, arr...)
		kk--
	}
	for nn > 0 {
		length *= nn
		nn--
	}
	for len(arr) > 0 {
		target := length / len(arr)
		length %= n
		if (length==0) {
			length+=target
		}
		res += string(arr[target]+'0')
		arr = append(arr[:target], arr[target+1:]...)
	}
	return res
}
