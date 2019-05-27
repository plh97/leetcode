# [572. Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/)

## 2019/05/27

### 题目 💗[easy]

Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

给定两个非空二叉树 s 和 t, 检查 t 是否是 s 的子树, 也就是说 s 或者 s的子树 和t完全一致, 返回true ,否则返回 false

Example 1:
Given tree s:

```bash
     3
    / \
   4   5
  / \
 1   2
```

Given tree t:

```basg
   4
  / \
 1   2
```

Return true, because t has the same structure and node values with a subtree of s.
Example 2:
Given tree s:

```bash
     3
    / \
   4   5
  / \
 1   2
    /
   0
```

Given tree t:

```bash
   4
  / \
 1   2
```

Return false.
