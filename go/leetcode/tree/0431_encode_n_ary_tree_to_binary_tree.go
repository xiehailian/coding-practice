package tree

//设计一个算法，可以将 N 叉树编码为二叉树，并能将该二叉树解码为原 N 叉树。一个 N 叉树是指每个节点都有不超过 N 个孩子节点的有根树。类似地，一个二叉树是指每个节点都有不超过 2 个孩子节点的有根树。你的编码 / 解码的算法的实现没有限制，你只需要保证一个 N 叉树可以编码为二叉树且该二叉树可以解码回原始 N 叉树即可。
//
//例如，你可以将下面的 3-叉 树以该种方式编码：

//
//注意，上面的方法仅仅是一个例子，可能可行也可能不可行。你没有必要遵循这种形式转化，你可以自己创造和实现不同的方法。
//
//注意：
//
//N 的范围在 [1, 1000]
//不要使用类成员 / 全局变量 / 静态变量来存储状态。你的编码和解码算法应是无状态的。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/encode-n-ary-tree-to-binary-tree

type NTreeNode struct {
	Val int
	Children []*NTreeNode
}

func encode(root *NTreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := &TreeNode{root.Val, nil, nil}
	if len(root.Children) != 0 {
		res.Left = encode(root.Children[0])
	}

	cur := res.Left
	for i := 1; i < len(root.Children); i++ {
		cur.Right = encode(root.Children[i])
		cur = cur.Right
	}
	return res
}

func decode(root *TreeNode) *NTreeNode {
	if root == nil {
		return nil
	}

	res := &NTreeNode{root.Val, []*NTreeNode{}}
	cur := root.Left
	for cur != nil {
		res.Children = append(res.Children, decode(cur))
		cur = cur.Right
	}
	return res
}