package findSubstring

import (
	"sort"
	"strings"
)

func findSubstring(s string, words []string) []int {
	// 自由组合 words 成一系列字符串
	if len(words) == 0 {
		return []int{}
	// } else if len(s) > 2000 {
	// 	return []int{}
	}
	res := []int{}
	ss := getRadom(len(words))
	wordArray := []string{}
	for i := range ss {
		tempString := ""
		for j := range ss[i] {
			tempString += words[ss[i][j]-1]
		}
		wordArray = append(wordArray, tempString)
	}
	for i := range wordArray {
		tempS := s
		for strings.LastIndex(tempS, wordArray[i]) > -1 {
			resInt := strings.LastIndex(tempS, wordArray[i])
			tempS = tempS[:len(tempS)-1]
			if !isRepeat(resInt, res) {
				res = append(res, resInt)
			}
		}
	}
	sort.Ints(res)
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

func getRadom(n int) [][]int {
	// if len(dp) >= s {
	// 	fmt.Println("with dp", dp[s-1])
	// 	return dp[s-1]
	// }
	if res, ok := dp[n]; ok {
		return res
	}
	if n == 2 {
		return [][]int{
			[]int{1, 2},
			[]int{2, 1},
		}
	}
	res := [][]int{}
	for i := 1; i <= n; i++ {
		_lowerRes := getRadom(n - 1)
		lowerRes := [][]int{}
		for i := range _lowerRes {
			var temp = make([]int, len(_lowerRes[i]))
			copy(temp, _lowerRes[i])
			lowerRes = append(lowerRes, temp)
		}
		for j := range lowerRes {
			for k := range lowerRes[j] {
				if lowerRes[j][k] == i {
					lowerRes[j][k] = n
				}
			}
			lowerRes[j] = append(lowerRes[j], i)
		}
		res = append(res, lowerRes...)
	}
	dp[n] = res
	return res
}

// n->2 [[2]]
// n->1 [[1,2],[2,1]]
// n=> 动态规划+缓存
// 复杂度n*n, 自下而上动态规划, 过于复杂,无法利用到动态规划的优势
// func getRadom(s int, ss [][]int) [][]int {
// 	// res X s
// 	if s == 0 {
// 		return ss
// 	} else if len(ss) == 0 {
// 		return getRadom(s-1, [][]int{[]int{s}})
// 	}

// 	// if res, ok := dp[s]; ok {
// 	// 	fmt.Println("with dp")
// 	// 	return res
// 	// }
// 	res := [][]int{}
// 	// 轮训几次
// 	for i := range ss {
// 		a := ss[i] // [1,2,3,4] ...
// 		if len(ss[i]) == 0 {
// 			res = append(res, []int{s})
// 		} else {
// 			for k := 0; k <= len(ss[i]); k++ {
// 				res = append(res, []int{})
// 				if k != 0 {
// 					l := a[:k]
// 					res[i*len(res[i])+k] = append(res[i*len(res[i])+k], l...)
// 				}
// 				res[i*len(res[i])+k] = append(res[i*len(res[i])+k], s)
// 				if k != len(res[i]) {
// 					r := a[k:]
// 					res[i*len(res[i])+k] = append(res[i*len(res[i])+k], r...)
// 				}
// 			}
// 		}
// 		ss[i] = res[i]
// 	}
// 	// dp[s] = res
// 	return getRadom(s-1, res)
// }
