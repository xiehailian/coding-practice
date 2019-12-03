package design

type lruNode struct {
	prev *lruNode
	next *lruNode

	key   int
	value int}


type LRUCache struct {
	head *lruNode
	tail *lruNode

	m        map[int]*lruNode
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	head := &lruNode{nil, nil, -1, -1}
	tail := &lruNode{nil, nil, -1, -1}
	head.next = tail
	tail.prev = head
	m := make(map[int]*lruNode, capacity)
	return LRUCache{head, tail, m, capacity}
}

func (this *LRUCache) insert(node *lruNode) {
	tail := this.tail
	tail.prev.next = node
	node.prev = tail.prev
	node.next = tail
	tail.prev = node
}

func (this *LRUCache) remove(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.insert(node)
		node.value = value
	} else {
		if len(this.m) == this.capacity {
			tmp := this.head.next
			this.remove(tmp)
			delete(this.m, tmp.key)
		}
		node := &lruNode{nil, nil, key, value}
		this.m[key] = node
		this.insert(node)
	}
}