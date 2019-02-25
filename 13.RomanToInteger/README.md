## [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

###### 2019/02/24


### 题目💗
罗马字转数字
先定义出罗马字映射字符串列表
```go
var RomanMap [][]string = [][]string{
	[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	[]string{"", "M", "MM", "MMM", "CD", "D", "DC", "DCC", "DCCC", "DCCCC"},
}
```
那么剩余的问题就是将`12312`数字分割成`1`,`2`,`3`,`4`并一一映射