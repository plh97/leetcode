package duplicateZeros

func duplicateZeros(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == 0 {
			for j := length - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[i] = 0
			i++
		}
	}
	return arr
}
