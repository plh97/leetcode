# [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

## 2019/05/19

### 题目 💗[easy]

Given a binary tree, determine if it is height-balanced.

给定一个二叉树,确定他是否平衡

---

For this problem, a height-balanced binary tree is defined as:

对于这个问题, 这个平衡二叉树如下定义

---

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

一个二叉树是每一个根节点到 root 顶点的高度差: `深度差 <= 1`

---

Example 1:

例子 1:

---

```bash
Given the following tree [3,9,20,null,null,15,7]:
    3
   / \
  9  20
    /  \
   15   7
Return true.
```

---

Example 2:

例子 2:

---

```bash
Given the following tree [1,2,2,3,3,null,null,4,4]:
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
```
