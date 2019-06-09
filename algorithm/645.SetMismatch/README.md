# [645. Set Mismatch](https://leetcode.com/problems/set-mismatch/)

## 2019/06/10

### 题目 💗[easy]

The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.

给定一个 S 集合,包括数字从 1-n, 但是不幸的是这个集合是乱序的,我们因为数据错误,一个集合中的数字重复了,导致原本的数字丢失.

---

Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

给定一个数组集合,`[重复的数字, 漏掉的数字]`

---

Example 1:

```b
Input: nums = [1,2,2,4]
Output: [2,3]
```

Note:

1. The given array size will in the range [2, 10000].
   给定的数组大小在 2-10000 之间
2. The given array's numbers won't have any order.
   这个给定的数组将不会拥有另一个顺序
