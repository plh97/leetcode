# [628. Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/)

## 2019/05/01

### 推荐程度

👍642 👎253

### 题目 💗[easy]

Given an integer array, find three numbers whose product is maximum and output the maximum product.

---

给定数组, 找出三个数相乘能得出最大结果

### 假定结果

先排序

[1,2,3,4,5,6]

1. 情况 `4*5*6`

2. 情况 `6*1*2` 基于 1,2 同时为负数的情况

`Math.max(a,b)`