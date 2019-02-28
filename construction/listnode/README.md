### 单链表

### n叉树
```
n=1 -> 单链表
n=2 -> 二叉树
n=n -> 参照 html 树
```

### description
single list node next to next, 

### Single Node
```go
type ListNode struct {
  Val  int
  Next *ListNode
}
```