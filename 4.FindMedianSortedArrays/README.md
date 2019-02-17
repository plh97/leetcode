## [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/submissions/)

###### 2019/02/17

> #### 思路
> [1,3][-1,1] 先合并数组,然后排序,我用的冒泡排序 复杂度 2(n+m)
> 长度为偶数,返回 num3[len(num3)/2] + 返回 num3[len(num3)-1/2]
> ```go
> if len(nums3)%2 == 0 {
> 	return float64(nums3[(len(nums3)-1)/2]+nums3[(len(nums3))/2])/2
> } else {
> 	return float64(nums3[(len(nums3)-1)/2])
> }
> ```

### TODO
- [x] 测试用例通过