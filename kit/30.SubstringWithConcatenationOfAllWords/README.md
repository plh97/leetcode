## [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)

###### 2019/03/04


### 题目💗 [hard]
这个应该不是什么好题目,应该放弃...





### 功能函数
给定 `n` 如何获取`1~n`有多少种排列可能 事实上,如果不用动态规划来做一层缓存,n=7我的内存就炸了
```
1 => n
2 => 12,21
3 => 123,132,213,231,312,321 (6种)
4 => 利用 `n=3`✖️4 => 24种可能
5 => 24*5 120种排列可能
6 => 120*6 720 种可能
7 => 720*7 = 5040 种可能
n => n!种可能
```

不缓存,只利用分解小问题的做法
```go
// n=> 动态规划来做啊,不然肯定卡死啊
func getRadom(s int, ss []string) []string {
	// res X s
	if s == 0 {
		return ss
	}
	res := []string{}
	for i := range ss {
		if len(ss[i]) == 0 {
			res = append(res, strconv.Itoa(s))
		} else {
			for k := 0; k <= len(ss[i]); k++ {
				l := ss[i][:k]
				r := ss[i][k:]
				res = append(res, l+strconv.Itoa(s)+r)
			}
		}
	}
	return getRadom(s-1, res)
}
```