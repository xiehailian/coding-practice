package linkedlist

//编写一个程序，找到两个单链表相交的起始节点。
//
//如下面的两个链表：
//
//在节点 c1 开始相交。
//
//示例 1：
//
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	ha := headA
	hb := headB
	for ha != hb {
		if ha == nil {
			ha = headB
		} else {
			ha = ha.Next
		}
		if hb == nil {
			hb = headA
		} else {
			hb = hb.Next
		}
	}
	return ha
}
