# [1005. Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/)

## 2019/06/05

### 题目 💗[easy]

Given an array A of integers, we must modify the array in the following way: we choose an i and replace A[i] with -A[i], and we repeat this process K times in total. (We may choose the same index i multiple times.)

给定一个整数数组,我们必须修改他k次,就是将数组内的任意数字取反.我们也可以将同一个数字取反n次.

Return the largest possible sum of the array after modifying it in this way.

Example 1:

```bash
Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].
```

Example 2:

```bash
Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
```

Example 3:

```bash
Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
```

Note:

1. 1 <= A.length <= 10000
2. 1 <= K <= 10000
3. -100 <= A[i] <= 100
