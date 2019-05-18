# [728. Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/)

## 2019/05/11

### 题目 💗[easy]

返回不可自整除的数字.

A self-dividing number is a number that is divisible by every digit it contains.

一个可自我整除的数字----这个数字是可以被他从每一位数(百位数,十位数,个位数)整除.

---

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

举个例子, 128 是可自我整除的数字, 因为 `128/(1|2|8)`.

--

Also, a self-dividing number is not allowed to contain the digit zero.

同时,一个可自我整除的数字,不允许 0 存在

---

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

给一个较低的数字 - 较高的范围, 输出一系列可自整除的数字,包括边界数字,如果可能的话.
