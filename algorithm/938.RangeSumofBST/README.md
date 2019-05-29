# [938. Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/)

## 2019/05/29

### 题目 💗[easy]

Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

给定一个根节点二叉查找树,返回 L 和 r 的总合.

---

The binary search tree is guaranteed to have unique values.

二叉查找树被保证拥有唯一值.

---

Example 1:

> Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
> Output: 32

Example 2:

> Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
> Output: 23

Note:

- 1.The number of nodes in the tree is at most 10000. 1.这个节点的总数<10000
- 2.The final answer is guaranteed to be less than 2^31. 2.这个最终答案保证小于 1<<31
