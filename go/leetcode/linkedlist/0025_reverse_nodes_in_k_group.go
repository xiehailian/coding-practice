package linkedlist

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//示例 :
//
//给定这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
//说明 :
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group

func reverseKGroup(head *ListNode, k int) *ListNode {

	var dummy = &ListNode{Val: 0}
	dummy.Next = head

	cur, end := dummy, dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := cur.Next
		next := end.Next
		end.Next = nil
		cur.Next = reverse(start)
		start.Next = next
		cur = start
		end = cur
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
