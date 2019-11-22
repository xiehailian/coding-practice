package linkedlist

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs

func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		start := cur.Next
		end := cur.Next.Next
		cur.Next = end
		start.Next = end.Next
		end.Next = start
		cur = start
	}
	return dummy.Next
}