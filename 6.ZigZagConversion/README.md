## [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

###### 2019/02/19

> #### 思路
> 分三种情况
> #### 第一行
> for 循环一次
> #### 中间行
> 两层for循环
> #### 最后一行
> 也就是第一行+nowRows-2


### TODO
- [x] 测试用例通过
- [x] 一层for循环遍历一般都很慢,应该尽量避免写O(m*n)代码