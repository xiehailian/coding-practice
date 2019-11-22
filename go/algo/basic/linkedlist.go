package basic

import (
	"errors"
)

type Node struct {
	value interface{}
	next *Node
}

func NewNode(value interface{}, next *Node) *Node {
	return &Node{value, next}
}

type LinkedList struct {
	dummy *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{&Node{nil, nil}, 0}
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) isEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Add(index int, value interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("add failed")
	}

	prev := l.dummy
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	cur := prev.next
	prev.next = NewNode(value, cur)
	l.size++
	return nil
}

func (l *LinkedList) AddFirst(value interface{}) error {
	return l.Add(0, value)
}

func (l *LinkedList) AddLast(value interface{})  error {
	return l.Add(l.size, value)
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("get failed")
	}

	cur := l.dummy.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.value, nil
}

func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}


