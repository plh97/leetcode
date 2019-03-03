package maxArea

func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1

	for i < j {
		ii, jj := height[i], height[j]
		tempMax := min(ii, jj) * (j - i)
		if tempMax > max {
			max = tempMax
		}
		if ii > jj {
			j--
		} else {
			i++
		}
	}
	return max
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
