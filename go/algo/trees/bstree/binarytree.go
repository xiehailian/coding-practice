package bstree

import (
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(v int) *Node { return &Node{value: v} }

type BST struct {
	root *Node
	size int
}

func NewBST() *BST { return &BST{root: nil, size: 0} }

func (b *BST) Add(v int) {
	if b.root == nil {
		b.root = NewNode(v)
		b.size++
	} else {
		b.add(b.root, v)
	}
}

func (b *BST) add(n *Node, v int) {
	if n.value == v {
		return
	} else if n.value < v && n.right == nil {
		n.right = NewNode(v)
		b.size++
	} else if n.value > v && n.left == nil {
		n.left = NewNode(v)
		b.size++
	}

	if n.value < v {
		b.add(n.right, v)
	} else {
		b.add(n.left, v)
	}
}

func (b *BST) Contains(v int) bool {
	return b.contains(b.root, v)
}

func (b *BST) contains(n *Node, v int) bool {
	if n == nil {
		return false
	}

	if n.value == v {
		return true
	} else if n.value > v {
		return b.contains(n.left, v)
	} else if n.value < v {
		return b.contains(n.right, v)
	}
	return false
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	b.preOrder(n.left)
	b.preOrder(n.right)
}

func (b *BST) PreOrderNR() {
	if b.root == nil {
		return
	}

	stack := []*Node{}
	stack = append(stack, b.root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.value)

		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

func (b *BST) InOrder() {
	b.inOrder(b.root)
}

func (b *BST) inOrder(n *Node) {
	if n == nil {
		return
	}
	b.inOrder(n.left)
	fmt.Println(n.value)
	b.inOrder(n.right)
}

func (b *BST) InOrderNR() {
	if b.root == nil {
		return
	}

	stack := []*Node{}
	root := b.root
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		if len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(cur.value)
			root = cur.right
		}
	}
}

func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

func (b *BST) postOrder(n *Node) {
	if n == nil {
		return
	}

	b.postOrder(n.left)
	b.postOrder(n.right)
	fmt.Println(n.value)
}

func (b *BST) PostOrderNR() {
	if b.root == nil {
		return
	}

	var pre *Node = nil
	cur := b.root
	stack := []*Node{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur.right == nil || pre == cur.right {
				fmt.Println(cur.value)
				pre = cur
				cur = nil
			} else {
				stack = append(stack, cur)
				cur = cur.right
			}
		}
	}
}

func (b *BST) LevelOrder() {
	if b.root == nil {
		return
	}

	list := list.New()
	list.PushFront(b.root)
	for list.Len() > 0 {
		cur := list.Remove(list.Back()).(*Node)
		fmt.Println(cur.value)

		if cur.left != nil {
			list.PushFront(cur.left)
		}
		if cur.right != nil {
			list.PushFront(cur.right)
		}
	}
}

func (b *BST) RemoveMin() int {
	res := b.Minimum()
	b.root = b.removeMin(b.root)
	return res
}

func (b *BST) Minimum() int {
	min := b.minimum(b.root)
	return min.value
}

func (b *BST) minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return b.minimum(n.left)
}

func (b *BST) removeMin(n *Node) *Node {
	if n.left == nil {
		r := n.right
		n.right = nil
		b.size--
		return r
	}

	n.left = b.removeMin(n.left)
	return n
}

func (b *BST) RemoveMax() int {
	res := b.Maximum()
	b.root = b.removeMax(b.root)
	return res
}

func (b *BST) Maximum() int {
	max := b.maximum(b.root)
	return max.value
}

func (b *BST) maximum(n *Node) *Node {
	if n.right == nil {
		return n
	}
	return b.maximum(n.right)
}

func (b *BST) removeMax(n *Node) *Node {
	if n.right == nil {
		l := n.left
		n.left = nil
		b.size--
		return l
	}

	n.right = b.removeMax(n.right)
	return n
}

func (b *BST) Remove(value int) {
	b.root = b.remove(b.root, value)
}

func (b *BST) remove(n *Node, value int) *Node {

	if n == nil {
		return nil
	}

	if n.value > value {
		n.left = b.remove(n.left, value)
		return n
	} else if n.value < value {
		n.right = b.remove(n.right, value)
		return n
	} else {
		if n.left == nil {
			r := n.right
			n.right = nil
			b.size--
			return r
		}

		if n.right == nil {
			l := n.left
			n.left = nil
			b.size--
			return l
		}

		s := b.minimum(n.right)
		s.right = b.removeMin(n.right)
		s.left = n.left
		n.left, n.right = nil, nil
		return s
	}
}

func (b *BST) Print() {
	var res = ""
	b.generateBSTString(b.root, 0, &res)
	fmt.Println(res)
}

func (b *BST) generateBSTString(n *Node, depth int, res *string) {
	if n == nil {
		*res = *res + b.generateDepthString(depth) + "nil\n"
		return
	}

	*res = *res + b.generateDepthString(depth) + strconv.Itoa(n.value) + "\n"
	b.generateBSTString(n.left, depth+1, res)
	b.generateBSTString(n.right, depth+1, res)
}

func (b *BST) generateDepthString(depth int) string {
	var res string
	for i := 0; i < depth; i++ {
		res = res + "--"
	}
	return res
}
