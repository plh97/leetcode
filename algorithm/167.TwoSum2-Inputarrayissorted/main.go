package twoSum

func twoSum(numbers []int, target int) []int {
	res := []int{}
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{
					i + 1, j + 1,
				}
			} else if numbers[j] > target && numbers[i] < target {
				i++
				j = i
			} else if numbers[j] > target && numbers[i] > target {
				return []int{}
			}
		}
	}

	return res
}
