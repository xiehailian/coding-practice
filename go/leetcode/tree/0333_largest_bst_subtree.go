package tree

//给定一个二叉树，找到其中最大的二叉搜索树（BST）子树，其中最大指的是子树节点数最多的。
//
//注意:
//子树必须包含其所有后代。
//
//示例:
//
//输入: [10,5,15,1,8,null,7]
//
//   10
//   / \
//  5  15
// / \   \
//1   8   7
//
//输出: 3
//解释: 高亮部分为最大的 BST 子树。
//     返回值 3 在这个样例中为子树大小。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/largest-bst-subtree

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isValidBST(root) {
		return getCount(root)
	}

	left := largestBSTSubtree(root.Left)
	right := largestBSTSubtree(root.Right)
	if left > right {
		return left
	} else {
		return right
	}
}

func getCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getCount(root.Left) + getCount(root.Right)
}