package design

import (
	"sync"
)

type BlockingQueueNode struct {
	value int
	prev  *BlockingQueueNode
	next  *BlockingQueueNode
}

type BoundedBlockingQueue struct {
	head *BlockingQueueNode
	tail *BlockingQueueNode

	enqLock sync.Mutex
	deqLock sync.Mutex

	enqBlock      chan int
	deqBlock      chan int
	enqBlockState bool
	deqBlockState bool

	size     int
	capacity int
}

func NewBoundedBlockingQueue(capacity int) BoundedBlockingQueue {
	head := &BlockingQueueNode{0, nil, nil}
	tail := &BlockingQueueNode{0, nil, nil}
	head.next = tail
	tail.prev = head
	return BoundedBlockingQueue{
		head:     head,
		tail:     tail,
		enqBlock: make(chan int),
		deqBlock: make(chan int),
		size:     0,
		capacity: capacity,
	}
}

func (this *BoundedBlockingQueue) insert(node *BlockingQueueNode) {
	tail := this.tail
	tail.prev.next = node
	node.prev = tail.prev
	node.next = tail
	tail.prev = node
}

func (this *BoundedBlockingQueue) remove(node *BlockingQueueNode) {
	this.head.next = node.next
	node.next.prev = this.head
}

func (this *BoundedBlockingQueue) Enqueue(element int) {
	this.enqLock.Lock()
	defer this.enqLock.Unlock()

	if this.size >= this.capacity {
		this.enqBlockState = true
		this.enqBlock <- 1
		this.enqBlockState = false
	}

	node := &BlockingQueueNode{element, nil, nil}
	this.insert(node)
	this.size++

	if this.deqBlockState {
		<-this.deqBlock
	}
}

func (this *BoundedBlockingQueue) Dequeue() int {
	this.deqLock.Lock()
	defer this.deqLock.Unlock()

	if this.size == 0 {
		this.deqBlockState = true
		this.deqBlock <- 1
		this.deqBlockState = false
	}

	node := this.head.next
	this.remove(node)
	this.size--

	if this.enqBlockState {
		<-this.enqBlock
	}
	return node.value
}

func (this *BoundedBlockingQueue) Size() int {
	return this.size
}
