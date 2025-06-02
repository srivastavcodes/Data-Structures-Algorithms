package lru_cache

type Node struct {
	prev, next *Node
	key, val   int
}

type LRUCache struct {
	capacity    int
	cache       map[int]*Node
	left, right *Node
}

func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		left:     &Node{},
		right:    &Node{},
	}
	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (lru *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) insert(node *Node) {
	prev, next := lru.right.prev, lru.right

	prev.next = node
	next.prev = node

	node.next = next
	node.prev = prev
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.remove(node)
		lru.insert(node)
		return node.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.remove(node)
		delete(lru.cache, key)
	}
	node := &Node{key: key, val: value}
	lru.cache[key] = node
	lru.insert(node)

	if len(lru.cache) > lru.capacity {
		lrc := lru.left.next
		lru.remove(lrc)
		delete(lru.cache, lrc.key)
	}
}
