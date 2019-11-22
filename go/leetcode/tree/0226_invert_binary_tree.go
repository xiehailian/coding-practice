package tree

//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/invert-binary-tree

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var tmp *TreeNode
	if root.Left != nil || root.Right != nil {
		tmp = root.Right
		root.Right = root.Left
		root.Left = tmp
	}

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}