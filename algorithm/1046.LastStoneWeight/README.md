# [1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

## 2019/06/04

### 题目 💗[easy]

We have a collection of rocks, each rock has a positive integer weight.

我们收集石头.每个是否都是数组里面的正数

Each turn, we choose the two heaviest rocks and smash them together. Suppose the stones have weights x and y with x <= y. The result of this smash is:

每个回合,我们找到两个最大的石头.让他们相减,将答案推入数组.

If x == y, both stones are totally destroyed;

如果两个石头相同重量,他们双双互毁.

If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left. Return the weight of this stone (or 0 if there are no stones left.)

Example 1:

```bash
Input: [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
```

Note:

1. 1 <= stones.length <= 30
2. 1 <= stones[i] <= 1000
