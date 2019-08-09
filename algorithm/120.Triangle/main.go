package minimumTotal

// func minimumTotal(triangle [][]int) int {
// 	depth := len(triangle)
// 	if depth == 0 {
// 		return 0
// 	} else if depth == 1 {
// 		return triangle[0][0]
// 	}
// 	for i := depth - 2; i > -1; i-- {
// 		for j := 0; j < len(triangle[i]); j++ {
// 			if triangle[i+1][j] < triangle[i+1][j+1] {
// 				triangle[i][j] += triangle[i+1][j]
// 			} else {
// 				triangle[i][j] += triangle[i+1][j+1]
// 			}
// 		}
// 	}
// 	return triangle[0][0]
// }

func minimumTotal(triangle [][]int) int {
	return helper(triangle, 0, 0)
}

func helper(triangle [][]int, index, depth int) int {
	// 到底不再递归
	if depth == len(triangle) {
		return 0
	}
	left := helper(triangle, index, depth+1)
	right := helper(triangle, index+1, depth+1)
	return triangle[depth][index] + min(left, right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
