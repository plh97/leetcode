package capacity
type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type cache interface {
	Get(key int) int
	Put(key int, value int)
}

// LRUCache is a Least Recently Used (LRU) cache
// data structure made up of a doubly-linked list
// and a hash table.
type LRUCache struct {
	cache
	Capacity int
	Map      map[int]*node
	Head     *node
	Tail     *node
}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*node, capacity),
		Head:     nil,
		Tail:     nil,
	}
}

// Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
func (c *LRUCache) Get(key int) int {
	if found, ok := c.Map[key]; ok {
		c.remove(found)
		c.add(found)
		return found.val
	}
	return -1
}

func (c *LRUCache) remove(n *node) {
	if n == c.Head {
		c.Head = n.next
	}

	if n == c.Tail {
		c.Tail = n.prev
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
}

func (c *LRUCache) add(n *node) {
	n.prev = nil
	n.next = c.Head
	if n.next != nil {
		n.next.prev = n
	}
	c.Head = n
	if c.Tail == nil {
		c.Tail = n
	}
}

// Put the value if the key doesn't exist, otherwise remove LRU at capacity.
func (c *LRUCache) Put(key int, value int) {
	found, ok := c.Map[key]
	if !ok {
		found = &node{key: key, val: value}
		c.Map[key] = found
	} else {
		found.val = value
		c.remove(found)
	}

	c.add(found)
	if len(c.Map) > c.Capacity {
		found = c.Tail
		if found != nil {
			c.remove(found)
			delete(c.Map, found.key)
		}
	}
}
