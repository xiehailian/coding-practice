package basic

// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。

func FindInPartiallySortedMatrix(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	columns := len(matrix[0])

	for r, c := 0, columns-1; r < rows && c >= 0 ;{
		if matrix[r][c] == target {
			return true
		}

		if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}