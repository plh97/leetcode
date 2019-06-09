# [747. Largest Number At Least Twice of Others](https://leetcode.com/problems/largest-number-at-least-twice-of-others/)

## 2019/06/10

### 题目 💗[easy]

In a given integer array nums, there is always exactly one largest element.

给定一个整数数组,这里总有一个最大元素

---

Find whether the largest element in the array is at least twice as much as every other number in the array.

找出最大元素

---

If it is, return the index of the largest element, otherwise return -1.

如果这个最大元素是里面是里面其他元素的 2 倍以上.返回最大元素,否则返回-1

---

Example 1:

```bash
Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
```

Example 2:

```bash
Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
```

Note:

1. nums will have a length in the range [1, 50].
2. Every nums[i] will be an integer in the range [0, 99].
