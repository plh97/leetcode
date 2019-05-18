package isRobotBounded

func isRobotBounded(instructions string) bool {
	instructions += instructions
	instructions += instructions
	start := [2]int{0, 0}
	direct := [2]int{0, 1}

	allLeftDirect := [][2]int{
		[2]int{0, 1},
		[2]int{-1, 0},
		[2]int{0, -1},
		[2]int{1, 0},
	}
	for i := range instructions {
		if instructions[i] == 'G' {
			start[0] += direct[0]
			start[1] += direct[1]
		} else if instructions[i] == 'L' {
			// 左转, 不动
			index := indexOfTwoArray(allLeftDirect, direct)
			if index == 3 {
				direct = allLeftDirect[0]
			} else {
				direct = allLeftDirect[index+1]
			}
		} else if instructions[i] == 'R' {
			// 右转, 不动
			index := indexOfTwoArray(allLeftDirect, direct)
			if index == 0 {
				direct = allLeftDirect[3]
			} else {
				direct = allLeftDirect[index-1]
			}
		}
	}
	return start[0] == 0 && start[1] == 0
}

func indexOfTwoArray(nums [][2]int, b [2]int) int {
	res := -1
	for i := range nums {
		num := nums[i]
		if num[0] == b[0] &&
			num[1] == b[1] {
			res = i
		}
	}
	return res
}
