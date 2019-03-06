## [18. 4Sum](https://leetcode.com/problems/4sum/)

###### 2019/3/2


### 题目💗
给定n个数字,其中四个之和最接近 `taget` 数组

> ### 穷举法 计算超时/勉强通过
> ```go
> func fourSum(nums []int, target int) [][]int {
>   l := len(nums)
>   res := [][]int{}
>   for i := 0; i < l; i++ {
>     for j := i + 1; j < l; j++ {
>       for k := j + 1; k < l; k++ {
>         for m := k + 1; m < l; m++ {
>           currentArray := []int{nums[i], nums[j], nums[k], nums[m]}
>           sort.Ints(currentArray)
>           if nums[i]+nums[j]+nums[k]+nums[m] == target && isRepeat(res, currentArray) {
>             res = append(res, currentArray)
>           }
>         }
>       }
>     }
>   }
>   return res
> }
> 
> func isRepeat(arrs [][]int, arr []int) bool {
>   for i := range arrs {
>     isSame := true
>     for j := 0; j < len(arr); j++ {
>       if arr[j] != arrs[i][j] {
>         isSame = false
>       }
>     }
>     if isSame {
>       return false
>     }
>   }
>   return true
> }
> ```

