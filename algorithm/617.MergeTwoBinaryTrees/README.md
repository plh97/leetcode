# [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)

## 2019/05/28

### 推荐程度

👍1700 👎126

### 题目 💗[easy]

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

给两个二叉树, 合并他们

---

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

你需要合并他们成一个新的二叉树, 这个合并规则就是如果两个节点都存在,就相加.只有一个存在就复制一个节点.否咋,节点为空.

---

Example 1:

Input:

```bash
  Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
```

Output:

```bash
Merged tree:
       3
      / \
     4   5
    / \   \
   5   4   7
```

Note: The merging process must start from the root nodes of both trees.
