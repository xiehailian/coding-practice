package linkedlist

//给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
//
//要求返回这个链表的深拷贝。 
//
//示例：
//
//输入：
//{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
//
//解释：
//节点 1 的值是 1，它的下一个指针和随机指针都指向节点 2 。
//节点 2 的值是 2，它的下一个指针指向 null，随机指针指向它自己。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	m := make(map[*Node]*Node, 0)
	for cur != nil {
		clone := &Node{cur.Val, nil, nil}
		m[cur] = clone
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]

}