package twoSum

// 解题思路,运用hash表格特性,分别用key/value记录a,b的值
// m  => [a的index, b的值]
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := m[num]; ok {
			return []int{idx, i}
		}
		m[target-num] = i
	}
	return nil
}
