package tree


//给你一个二叉搜索树和其中的某一个结点，请你找出该结点在树中顺序后继的节点。
//
//结点 p 的后继是值比 p.val 大的结点中键值最小的结点。
//
// 
//
//示例 1:
//
//
//
//输入: root = [2,1,3], p = 1
//输出: 2
//解析: 这里 1 的顺序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/inorder-successor-in-bst

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	var res *TreeNode
	if root != nil {
		if root.Val > p.Val {
			res := inorderSuccessor(root.Left, p)
			if res != nil {
				return res
			}
			return root
		}
		return inorderSuccessor(root.Right, p)
	}
	return res
}

func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	for root != nil {
		if root.Val > p.Val {
			res, root = root, root.Left
		} else {
			root = root.Right
		}
	}
	return res
}

