# [541. Reverse String II](https://leetcode.com/problems/reverse-string-ii/)

## 2019/06/09

### 题目 💗[easy]

Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.

给定一个字符串和一个整数 K, 你需要反正第一个 k-2k 的字符串,如果这里的单词少于 k 个,反正所有的,如果这里少于 2k 但是等于 k 个字符,反正第一个 k 字符.

---

Example:

```bash
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
```

Restrictions:

1. The string consists of lower English letters only.
2. Length of the given string and k will in the range [1, 10000]
