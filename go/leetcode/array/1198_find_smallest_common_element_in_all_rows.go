package array

//给你一个矩阵 mat，其中每一行的元素都已经按 递增 顺序排好了。请你帮忙找出在所有这些行中 最小的公共元素。
//
//如果矩阵中没有这样的公共元素，就请返回 -1。
//
//
//
//示例：
//
//输入：mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
//输出：5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-smallest-common-element-in-all-rows

func smallestCommonElement(mat [][]int) int {

	if mat == nil || len(mat) == 0 {
		return -1
	}

	var row, column = len(mat), len(mat[0])
	for i := 0; i < column; i++ {
		for j := 1; j < row; j++ {
			left, right := 0, column-1
			for left < right {
				mid := left + (right-left)/2
				if mat[j][mid] < mat[0][i] {
					left = mid + 1
				} else {
					right = mid
				}
			}

			if mat[j][left] == mat[0][i] {
				continue
			} else {
				goto b
			}
		}
		return mat[0][i]
	b:
	}
	return -1
}
