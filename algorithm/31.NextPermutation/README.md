# [31. Next Permutation](https://leetcode.com/problems/next-permutation/)

## 2019/03/23

## 题目 💗[Medium]

1. 前后两个数对比,如果呈现递增关系,那就标记这个数的 index 为 i
2. 顺着`index=i`往后走,遇到第一个比他大 1 的数`index=j`, `swap(arr[i],arr[j])`
3. `reverse(arr[i+1:])`
