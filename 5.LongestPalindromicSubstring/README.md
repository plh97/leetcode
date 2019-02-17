## [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

###### 2019/02/17

> #### 思路
> abcbd
> 二次循环
> a
> ab
> abc
> abcb
> abcbd
> b
> bc
> bcb -> 检测是否是是回文函数


### TODO
- [x] 测试用例通过
- [ ] 双层for循环遍历一般都很慢,应该尽量避免写O(m*n)代码