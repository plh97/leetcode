package judgeCircle

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, e := range moves {
		switch e {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}

// func judgeCircle(moves string) bool {
// 	Map := make(map[rune]int, len(moves))
// 	for _, e := range moves {
// 		Map[e]++
// 	}
// 	if Map['L'] != Map['R'] || Map['U'] != Map['D'] {
// 		return false
// 	}
// 	return true
// }
