package search

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		i := (end + start) / 2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			start = i + 1
		} else {
			end = i - 1
		}
	}
	return -1
}

// func search(nums []int, target int) int {
// 	for i := range nums {
// 		if nums[i] == target {
// 			return i
// 		}
// 	}
// 	return 0
// }
