# [235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

## 2019/05/20

### 题目 💗[easy]

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

给定一个二叉树,找出两个元素的共通祖先

---

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

根据`LCA`的 wiki 定义,共通祖先就是共同祖先

---

Given binary search tree: root = [6,2,8,0,4,7,9,null,null,3,5]

Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Note:

注意:

---

All of the nodes' values will be unique.

所有节点的 value 是序列.

---

p and q are different and both values will exist in the BST.

p 和 q 是不同的并且都出现于 BST
