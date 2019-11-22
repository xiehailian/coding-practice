package basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	// len和cap的区别
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}

	var arr = NewArray(100)
	for i := 0; i < 10; i++  {
		arr.Add(i, s[i])
	}

	arr.Print()

	arr.Add(5, 100)
	arr.AddFirst(-1)
	arr.AddFirst(0)
	arr.Print()

	arr.Remove(5)
	arr.RemoveFirst()
	arr.RemoveLast()
	arr.Set(1, 100)
	arr.Print()
}

func TestLinkedList(t *testing.T)  {
	l := NewLinkedList()
	e1 := l.PushFront(1)
	l.InsertAfter(2, e1)
	l.InsertAfter(3, e1.Next())
	l.InsertAfter(4, e1.Next().Next())
	l.InsertAfter(5, e1.Next().Next().Next())

	l.Print()
	fmt.Println()
}

