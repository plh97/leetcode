package findSubstring

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	// 自由组合 words 成一系列字符串
	// if len(words) == 0 {
	// 	return []int{}
	// }else if len(s) > 2000 {
	// 	return []int{}
	// }
	res := []int{}

	ss := getRadom(4, [][]int{})
	fmt.Println(ss)
	// for i := range ss {
	// 	targetString := ""
	// 	for k := range ss[i] {
	// 		i, _ := strconv.Atoi(ss[i][k : k+1])
	// 		targetString += words[i-1]
	// 	}
	// 	// 不断消耗,源字符串.重复检测
	// 	ss := s
	// 	for len(ss) > 0 {
	// 		index := strings.LastIndex(ss, targetString)
	// 		if index > -1 {
	// 			ss = ss[:len(ss)-1]
	// 			if !isRepeat(index, res) {
	// 				res = append(res, index)
	// 			}
	// 		} else {
	// 			ss = ""
	// 		}
	// 	}
	// }
	// sort.Ints(res)
	return res
}

func isRepeat(num int, arr []int) bool {
	for i := range arr {
		if arr[i] == num {
			return true
		}
	}
	return false
}

var dp map[int][][]int = make(map[int][][]int)

// n=> 动态规划+缓存来做啊,不然肯定卡死啊
func getRadom(s int, ss [][]int) [][]int {
	// res X s
	if s == 0 {
		return ss
	} else if len(ss) == 0 {
		return getRadom(s-1, [][]int{[]int{s}})
	}

	if res, ok := dp[s]; ok {
		return res
	}
	res := [][]int{}
	for i := range ss {
		a := ss[i] // [1,2,3,4] ...
		if len(ss[i]) == 0 {
			res = append(res, []int{s})
		} else {
			for k := 0; k <= len(ss[i]); k++ {
				fmt.Println(ss)
				l := a[:k]
				r := a[k:]
				res = append(res, l, []int{s}, r)
			}
		}
		ss[i] = res[i]
	}
	dp[s] = res
	return getRadom(s-1, res)
}

// 1,2,3,4

// []
// 1

// a = [
// 	[1,2,3],
// 	[1,3,2],
// 	[2,3,1],
// 	[2,1,3],
// 	[3,2,1],
// 	[3,1,2],
// ]
