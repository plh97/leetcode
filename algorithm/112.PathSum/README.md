# [112. Path Sum](https://leetcode.com/problems/path-sum/)

## 2019/05/20

### 题目 💗[easy]

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

给定一个二叉树综合, 检查是否存在某条 root-> bottom 的路径 value 相加 等于 给定值.

---

Note: A leaf is a node with no children.

注意: 一个`leaf`就是一个节点没有子节点.

Example:

Given the below binary tree and sum = 22,

```bash
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
```

return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

返回 true, 作为存在端到底部`5-4-11-2`

### 思路 leaf 循环遍历

