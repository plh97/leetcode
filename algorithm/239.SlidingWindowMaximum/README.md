# [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)

## 2019/05/21

### 题目 💗[easy]

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

给定一个数组. 这里有一个滑动窗口,宽度 k,从左移动到右边.返回由每个窗口中最大数字组成的数组.

---

Example:

```bash
Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

注意:
你可以假定有 k 总是有效的, 1<=k<=len(array).

---

Follow up:
Could you solve it in linear time?

下面:
你可以用 O(n)解决这个问题吗?
