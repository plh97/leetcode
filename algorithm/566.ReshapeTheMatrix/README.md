# [566. Reshape the Matrix](https://leetcode.com/problems/reshape-the-matrix/)

## 2019/05/01

### 推荐程度

👍550 👎83

### 题目 💗[easy]

In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

在`Matlab`中, 这里你有一个非常有用的函数叫做`reshape`,这个函数能够将一个矩阵转换成另一个新的矩阵,同时可以保存它的元数据

---

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

你被给一个拥有两个子数组的数组,并且还有整数`r`和整数`c`代表,它们分别排数和列数.

---

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

这个矩阵转换前后需要和以相同遍历次序.

---

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

如果给定条件`r`和`c`无法按照条件转换,则输出原矩阵.
