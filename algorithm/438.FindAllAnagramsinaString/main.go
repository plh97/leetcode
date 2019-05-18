package findAnagrams

import (
	"bytes"
	"strings"
)

func findAnagrams(s string, p string) []int {
	if len(s) > 20100 || len(p) > 20100 {
		return []int{}
	}
	res := []int{}
	for i, j := 0, len(p); j <= len(s); {
		str1 := s[i:j]
		str2 := copyString(p)
		if randomSameString(str1, str2) {
			res = append(res, i)
		}
		i++
		j++
	}
	return res
}

func randomSameString(str1, str2 string) bool {
	for i := range str1 {
		index := strings.Index(str2, str1[i:i+1])
		if index > -1 {
			str2 = str2[:index] + str2[index+1:]
		} else {
			return false
		}
	}
	return true
}

// 这种解法超时了
// func findAnagrams(s string, p string) []int {
// 	res := []int{}
// 	a := getRandom(len(p))
// 	// 遍历随机组合
// 	for i := range a {
// 		var str_buffer bytes.Buffer
// 		// 替换随机组合每个字符串
// 		for j := range a[i] {
// 			str_buffer.WriteString(p[a[i][j]-1 : a[i][j]])
// 		}
// 		str := str_buffer.String()
// 		for tempStr, index := copyString(s), strings.LastIndex(copyString(s), str); index > -1; {
// 			if index > -1 {
// 				index = strings.LastIndex(tempStr, str)
// 				if !duplicate(res, index) {
// 					res = append(res, index)
// 				}
// 				tempStr = tempStr[:len(tempStr)-1]
// 				index = strings.LastIndex(tempStr, str)
// 			}
// 		}
// 	}
// 	sort.Ints(res)
// 	return res
// }

// var dp map[int][][]int = make(map[int][][]int)

// func getRandom(n int) [][]int {
// 	if res, ok := dp[n]; ok {
// 		fmt.Println("get", res)
// 		return res
// 	}
// 	if n == 1 {
// 		return [][]int{
// 			[]int{1},
// 		}
// 	}
// 	if n == 2 {
// 		return [][]int{
// 			[]int{1, 2},
// 			[]int{2, 1},
// 		}
// 	}
// 	res := [][]int{}
// 	for i := 1; i <= n; i++ {
// 		_lowerRes := getRandom(n - 1)
// 		lowerRes := [][]int{}
// 		for i := range _lowerRes {
// 			var temp = make([]int, len(_lowerRes[i]))
// 			copy(temp, _lowerRes[i])
// 			lowerRes = append(lowerRes, temp)
// 		}
// 		for j := range lowerRes {
// 			for k := range lowerRes[j] {
// 				if lowerRes[j][k] == i {
// 					lowerRes[j][k] = n
// 				}
// 			}
// 			lowerRes[j] = append(lowerRes[j], i)
// 		}
// 		res = append(res, lowerRes...)
// 	}
// 	dp[n] = res
// 	return res
// }

func copyString(str string) string {
	var str_buffer bytes.Buffer
	for i := range str {
		str_buffer.WriteString(str[i : i+1])
	}
	return str_buffer.String()
}

// func duplicate(nums []int, num int) bool {
// 	for i := range nums {
// 		if nums[i] == num {
// 			return true
// 		}
// 	}
// 	return false
// }
