package design

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLruCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))

	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))

	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))

}

func TestLFUCache(t *testing.T) {
	cache := NewLFUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))

	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))
	assert.Equal(t, 3, cache.Get(3))

	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))
}

func TestMinStack(t *testing.T) {
	stack := NewMinStack()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(t, -3, stack.GetMin())

	stack.Pop()
	assert.Equal(t, 0, stack.Top())
	assert.Equal(t, -2, stack.GetMin())
}

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}
	iterator := NewBSTIterator(root)
	assert.Equal(t, 3, iterator.Next())
	assert.Equal(t, 7, iterator.Next())
	assert.Equal(t, true, iterator.HasNext())
	assert.Equal(t, 9, iterator.Next())
	assert.Equal(t, true, iterator.HasNext())
	assert.Equal(t, 15, iterator.Next())
	assert.Equal(t, true, iterator.HasNext())
	assert.Equal(t, 20, iterator.Next())
	assert.Equal(t, false, iterator.HasNext())
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, false, trie.Search("app"))
	assert.Equal(t, true, trie.StartsWith("app"))

	trie.Insert("app")
	assert.Equal(t, true, trie.Search("app"))
}

func TestSkiplist(t *testing.T) {
	skiplist := NewSkiplist()
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	assert.Equal(t, false, skiplist.Search(0))

	skiplist.Add(4)
	assert.Equal(t, true, skiplist.Search(1))
	assert.Equal(t, false, skiplist.Erase(0))
	assert.Equal(t, true, skiplist.Erase(1))
	assert.Equal(t, false, skiplist.Search(1))
}

func TestBoundedBlockingQueue(t *testing.T) {
	queue := NewBoundedBlockingQueue(3)
	go queue.Enqueue(1)
	go fmt.Println(queue.Dequeue())

	go queue.Dequeue()
	go queue.Enqueue(0)
	go queue.Enqueue(2)
	go queue.Enqueue(3)
	go queue.Enqueue(4)
	go queue.Dequeue()

	time.Sleep(1 * time.Second)
	fmt.Println(queue.Size())

}


func TestCircularDeque(t *testing.T) {
	queue := NewCircularDeque(3)
	assert.Equal(t, true, queue.InsertLast(1))
	assert.Equal(t, true, queue.InsertLast(2))
	assert.Equal(t, true, queue.InsertFront(3))
	assert.Equal(t, false, queue.InsertFront(4))
	assert.Equal(t, 2, queue.GetRear())
	assert.Equal(t, true, queue.IsFull())
	assert.Equal(t, true, queue.DeleteLast())
	assert.Equal(t, true, queue.InsertFront(4))
	assert.Equal(t, 4, queue.GetFront())
}