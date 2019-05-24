# [146. LRU Cache](https://leetcode.com/problems/lru-cache/)

## 2019/05/23

### 题目 💗[hard]

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

设计一个数据结构,至少能实现 (LRU) 缓存. 他应该支持接下来的操作, `get` 和 `put`

---

**get(key)** - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
**put(key, value)** - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

> **`get(key)`**
>
> 获取正数值, 如果它存在于缓存中, 如果不存在返回 `-1`,

---

> **`put(key,value)`**
>
> 如果这个键值不存在,设置或者插入值.当缓存触顶的它的容量,它应该被插入在最近的位置,使得最久,元素被推出.

---

Follow up:
Could you do both operations in O(1) time complexity?

你能用`O(n)` 的时间复杂度完成它吗?

---

Example:

例子:

---

```bash
LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```
