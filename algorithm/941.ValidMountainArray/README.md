# [941. Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/)

## 2019/05/03

### 题目 💗[easy]

Given an array A of integers, return true if and only if it is a valid mountain array.

给定两个整数数组, 根据它是否是山数组, 返回`true`/`false`,

---

Recall that A is a mountain array if and only if:

仅仅满足一下条件是山数组,

---

A.length >= 3

长度>3

---

There exists some i with 0 < i < A.length - 1 such that:

下面存在两种山数组情况

---

- `A[0] < A[1] < ... A[i-1] < A[i]`
- `A[i] > A[i+1] > ... > A[B.length - 1]`
