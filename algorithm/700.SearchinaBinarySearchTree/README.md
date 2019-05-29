# [700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)

## 2019/05/28

### 题目 💗[easy]

Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.
给定一个 BST 和一个值, 你需要找到节点的值与目标节点相同,找不到返回`nil`.

---

For example,
例如

---

Given the tree:

```bash
        4
       / \
      2   7
     / \
    1   3
```

And the value to search: 2

You should return this subtree:

```bash
      2
     / \
    1   3
```

In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL.
在上面的例子, 如果我们想要找到值 5 , 直到我们发现没有 5, 我们就返回 nil

---

Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as [], not null.
注意, 一个空的树代表 nil,
