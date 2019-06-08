# [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)

## 2019/06/09

### 题目 💗[easy]

In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
0 代表空格

---

the value 1 representing a fresh orange;
1 代表新鲜水果

---

the value 2 representing a rotten orange.
2 代表腐烂水果

---

Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

每分钟,都有一个腐烂水果的四周的新鲜水果腐烂

---

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1 instead.

返回所有水果腐烂的最小步数.

Example 1:

![images](oranges.png)

```bash
Input: [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
```

Example 2:

```b
Input: [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
```

Example 3:

```bash
Input: [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
```

Note:

1. 1 <= grid.length <= 10
2. 1 <= grid[0].length <= 10
3. grid[i][j] is only 0, 1, or 2.
