package randomMoney

import (
	"math/rand"
)

// 给100元, 随机生成10个红包,每个红包金额范围 6-12
//
// 分布比率 [12:151884 11:87687 8:40120 10:86157 6:32481 9:78519 7:34255]  这特么居然是个固定输出...
//
// 果然辣鸡算法
func randomMoney(total int) (int, []int) {
	less := []int{}
	more := []int{}

	for i := 0; i < 5; i++ {
		num1 := rand.Intn(6) + 6
		num2 := 20 - num1
		if num2 > 12 {
			less = append(less, num1)
			more = append(more, num2)
		} else {
			less = append(less, num1, num2)
		}
	}
	dx := 0

	for i, e := range more {
		if e > 12 {
			dx += e - 12
			more[i] = 12
		}
	}
	for i, e := range less {
		if dx == 0 {
			continue
		}
		if e+dx > 12 {
			dx -= 12 - e
			less[i] = 12
		} else {
			less[i] += dx
			dx = 0
		}
	}

	res := append(less, more...)
	sum := 0
	for _, e := range res {
		sum += e
	}
	return sum, append(less, more...)
}
