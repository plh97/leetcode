## [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

###### 2019/02/24


### 题目💗
罗马字转数字
先定义出罗马字映射字符串列表
```go
RomanMap := map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
```
`III`如何转成数字3呢,
这里采用消耗策略,`III`从左到右开始计算消耗,
`IV`如何转成数字4呢,
有一种例外就是`IV`,当每次循环的时候`I`<`V`的时候,将两个字母合并成 `res:=V-I`也就是4.
