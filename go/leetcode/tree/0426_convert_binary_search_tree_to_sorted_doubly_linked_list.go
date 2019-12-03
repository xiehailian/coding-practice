package tree

//将一个二叉搜索树就地转化为一个已排序的双向循环链表。可以将左右孩子指针作为双向循环链表的前驱和后继指针。
//
//为了让您更好地理解问题，以下面的二叉搜索树为例：
//
//
//
//我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
//
//下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
//
//
//特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。
//
//下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var cur, head *TreeNode
	inOrder(root, cur, head)
	cur.Right = head
	head.Left = cur
	return head
}

func inOrder(root, cur, head *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, cur, head)
	if head == nil {
		head = root
		cur = root
	} else {
		cur.Right = root
		root.Left = cur
		cur = root
	}
	inOrder(root.Right, cur, head)
}
