package tree

import (
	"strconv"
	"strings"
)

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//
//请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
//
//示例: 
//
//你可以将以下二叉树：
//
//    1
//   / \
//  2   3
//     / \
//    4   5
//
//序列化为 "[1,2,3,null,null,4,5]"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree

func serialize(root *TreeNode) string {
	return rserialize(root, "")
}

func rserialize(root *TreeNode, str string) string {
	if root == nil {
		return "null"
	} else {
		str = strconv.Itoa(root.Val)
		str += "," + rserialize(root.Left, str)
		str += "," + rserialize(root.Right, str)
	}
	return str
}

func deserialize(str string) *TreeNode {
	strs := strings.Split(str, ",")
	return rdserialize(strs)
}

func rdserialize(strs []string) *TreeNode {
	if strs[0] == "null" {
		copy(strs, strs[1:])
		return nil
	}

	value, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{value, nil, nil}
	copy(strs, strs[1:])
	root.Left = rdserialize(strs)
	root.Right = rdserialize(strs)

	return root
}