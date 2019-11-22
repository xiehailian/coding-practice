package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



func Test2(t *testing.T)  {
	tests := []struct{
		l1 *ListNode
		l2 *ListNode
		output *ListNode
	}{
		{
			l1: &ListNode{Val: 2, Next: &ListNode{Val:4, Next: &ListNode{Val: 3, Next: nil}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val:6, Next: &ListNode{Val: 4, Next: nil}}},
			output: &ListNode{Val: 7, Next: &ListNode{Val:0, Next: &ListNode{Val: 8, Next: nil}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := addTwoNumbers(tt.l1, tt.l2)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test19(t *testing.T) {
	tests := []struct{
		head *ListNode
		n int
		output *ListNode
	}{
		{
			head: &ListNode{Val: 2, Next: &ListNode{Val:4, Next: &ListNode{Val: 3, Next: nil}}},
			n: 2,
			output: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := removeNthFromEnd(tt.head, tt.n)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test23(t *testing.T)  {
	tests := []struct{
		input []*ListNode
		output *ListNode
	}{
		{
			input: []*ListNode{{Val: 1, Next: &ListNode{Val:4, Next: &ListNode{Val: 5, Next: nil}}}, {Val: 1, Next: &ListNode{Val:3, Next: &ListNode{Val: 4, Next: nil}}}, {Val: 2, Next: &ListNode{Val:6, Next: nil}}},
			output: &ListNode{Val: 1, Next: &ListNode{Val:1, Next: &ListNode{Val: 2, Next: &ListNode{Val:3, Next:&ListNode{Val:4, Next:&ListNode{Val:4, Next:nil}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := mergeKLists(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test24(t *testing.T)  {
	tests := []struct{
		input *ListNode
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			output: &ListNode{Val: 2, Next: &ListNode{Val:1, Next: &ListNode{Val: 4, Next: &ListNode{Val:3, Next: nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := swapPairs(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test25(t *testing.T)  {
	tests := []struct{
		input *ListNode
		k int
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			k: 2,
			output: &ListNode{Val: 2, Next: &ListNode{Val:1, Next: &ListNode{Val: 4, Next: &ListNode{Val:3, Next: nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := reverseKGroup(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test61(t *testing.T)  {
	tests := []struct{
		input *ListNode
		k int
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			k: 2,
			output: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: &ListNode{Val: 1, Next: &ListNode{Val:2, Next: nil}}}},
		},
		{
			input:  &ListNode{Val: 0, Next: &ListNode{Val:1, Next: &ListNode{Val: 2, Next: nil}}},
			k: 4,
			output: &ListNode{Val: 2, Next: &ListNode{Val:0, Next: &ListNode{Val: 1, Next: nil}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := rotateRight(tt.input, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test147(t *testing.T)  {
	tests := []struct{
		input *ListNode
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 4, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			output: &ListNode{Val: 2, Next: &ListNode{Val:3, Next: &ListNode{Val: 4, Next: &ListNode{Val:4, Next: nil}}}},
		},
		{
			input:  &ListNode{Val: 3, Next: &ListNode{Val:2, Next: &ListNode{Val: 1, Next: nil}}},
			output: &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: nil}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := insertionSortList(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test160(t *testing.T)  {
	tests := []struct{
		headA *ListNode
		headB *ListNode
		output *ListNode
	}{
		{
			headA:  &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			headB:  &ListNode{Val: 5, Next: &ListNode{Val:6, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			output: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := getIntersectionNode(tt.headA, tt.headB)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test328(t *testing.T)  {
	tests := []struct{
		input *ListNode
		output *ListNode
	}{
		{
			input:  &ListNode{Val: 4, Next: &ListNode{Val:2, Next: &ListNode{Val: 3, Next: &ListNode{Val:4, Next: nil}}}},
			output: &ListNode{Val: 4, Next: &ListNode{Val:3, Next: &ListNode{Val: 2, Next: &ListNode{Val:4, Next: nil}}}},
		},
		{
			input:  &ListNode{Val: 3, Next: &ListNode{Val:2, Next: &ListNode{Val: 1, Next: nil}}},
			output: &ListNode{Val: 3, Next: &ListNode{Val:1, Next: &ListNode{Val: 2, Next: nil}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := oddEvenList(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

