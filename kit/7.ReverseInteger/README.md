## [7. Reverse Integer](https://leetcode.com/problems/reverse-integer/)

###### 2019/02/19

> #### 思路
> ```go
> for x>0 循环
>   x -> 1234
>   res := 1
>   temp := x % 10 // 4
>   x = x / 10
>   res = res*10 + temp
> ```
> 这是最高效



### TODO
- [x] 测试用例通过
- [x] 双层for循环遍历一般都很慢,应该尽量避免写O(m*n)代码