## [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)

###### 2019/02/21

### 题目
数字是否是回文

> #### 思路
> ```go
> strconv.Atoi(str) // 核心函数 -> 返回数字以及错误
> strings.TrimSpace // 用于去除两端空格
> ```
> #### 特殊情况的考虑
> `-`,`+`,`.`
> 如果在去除空格后,首字符是`+`,`-`则continue,否则和字母,空格,.做同样处理,截取返回


### TODO
- [x] 测试用例通过
- [ ] 双层for循环遍历一般都很慢,应该尽量避免写O(m*n)代码