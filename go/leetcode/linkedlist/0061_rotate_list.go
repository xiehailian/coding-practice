package linkedlist

//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//示例 2:
//
//输入: 0->1->2->NULL, k = 4
//输出: 2->0->1->NULL
//解释:
//向右旋转 1 步: 2->0->1->NULL
//向右旋转 2 步: 1->2->0->NULL
//向右旋转 3 步: 0->1->2->NULL
//向右旋转 4 步: 2->0->1->NULL
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	old := head
	var count = 1
	for old.Next != nil {
		old = old.Next
		count++
	}

	old.Next = head

	newTail := head
	for i := 0; i < count-k%count-1; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	return newHead

}
