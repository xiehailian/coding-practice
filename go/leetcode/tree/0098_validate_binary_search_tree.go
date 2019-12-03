package tree

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//    2
//   / \
//  1   3
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/validate-binary-search-tree

func isValidBST(root *TreeNode) bool {
	return validBST(root, 0, 0)
}

func validBST(node *TreeNode, lower, higher int) bool {
	if node == nil {
		return true
	}

	var val = node.Val
	if lower != 0 && val <= lower {
		return false
	}
	if higher != 0 && val >= higher {
		return false
	}

	if !validBST(node.Right, val, higher) {
		return false
	}
	if !validBST(node.Left, lower, val) {
		return false
	}
	return true
}
