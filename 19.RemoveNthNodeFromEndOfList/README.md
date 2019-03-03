## [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

###### 2019/3/2

### 题目💗
链表, 删除倒数第n个节点
我觉得更换第`(len(n) - n)-1`那个节点的`Next`属性指针指向下下一个节点就好
过程中会遇到很多边界问题,需要自己写测试用例一个个覆盖,然后开启`golang debug`模式