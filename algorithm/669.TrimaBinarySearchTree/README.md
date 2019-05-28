# [669. Trim a Binary Search Tree](https://leetcode.com/problems/trim-a-binary-search-tree/)

## 2019/05/28

### 推荐程度

👍1119 👎129

### 题目 💗[easy]

Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

给定一个二叉树,并且给定另外两个数字 l 和 r,他们代表二叉树节点值的边界,如果节点值超出,则忽略这个节点,直接走下一个节点, 你可能需要改变树的根节点.所以结果应该返回新的二叉搜索树.

---

Example 1:
Input:

```bash
    1
   / \
  0   2

  L = 1
  R = 2
```

Output:

```bash
    1
      \
       2
```

Example 2:

例子 2:

Input:

```bash
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3
```

Output:

输出

---

```bash
      3
     /
   2
  /
 1
```
