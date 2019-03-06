###### 好吧,养成做笔记的习惯 2019/02/16

## [2.Add Two Numbers](https://leetcode.com/problems/add-two-numbers/submissions/)

> #### 思路
> 题目: `[1,2,3] + [9,9,9] -> [4,3,2,1]`
> 1.先在开头定义 最终返回数的 start 节点 以及 缓存节点
> ```go
> tempNode := &ListNode{}
> start := tempNode     // start 存着一个指针指向初始节点
> ```
> 2.这个传入的两个参数不是数组,这就需要我们换一种形式结题,不是以往的思路,比如循环数组.
> 如果当前 node 不是`nil`空指针的话,进行下一次迭代
> ```go
> for node != nil {
>   node.next = &ListNode{Val: (sum + add) % 10}
>   node = node.next                        
> }
> ```




