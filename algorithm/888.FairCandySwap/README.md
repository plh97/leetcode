# [888. Fair Candy Swap](https://leetcode.com/problems/fair-candy-swap/)

## 2019/05/02

### 题目 💗[easy]

Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.

alice 和 bob 有两种不同尺寸的糖,`A[i]`表示 alice 第 i 个糖的尺寸,同理`B[j]`

---

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy. (The total amount of candy a person has is the sum of the sizes of candy bars they have.)

直到他们成为朋友.他们决定交换一个糖,以便于互相拥有同等分量的糖

---

Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.

返回一个整数数组,是 alice 和 bob 交换的糖组成

---

If there are multiple answers, you may return any one of them. It is guaranteed an answer exists.

如果有重复的答案,你必须返回其中一个来保证答案存在.

## 答案

计算总量差值 n,循环找到相差 n/2 的数,调换
