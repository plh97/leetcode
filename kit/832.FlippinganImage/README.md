# [832. Flipping an Image](https://leetcode.com/problems/flipping-an-image/)

## 2019/05/02

### 题目 💗[easy]

Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

给定一个二进制矩阵 A, 我们想要翻转图片二进制,

---

To flip an image horizontally means that each row of the image is reversed. For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

为了`flip`图片,就是这样`[1,1,0]` => `[0,1,1]`, 其实就是翻转数组

---

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

为了`invert`一个图片意味着, 每个 `[0,1,1]` => `[1,0,0]`
