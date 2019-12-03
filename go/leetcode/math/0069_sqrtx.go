package math

//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//示例 2:
//
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sqrtx

func mySqrt(x int) int {

	if x == 0 {
		return 0
	}

	if x <= 3 {
		return 1
	}

	var l, r, m int64

	l = 0
	r = int64(x / 2)

	for l <= r {
		m = l + (r-l)/2
		if m*m == int64(x) {
			return int(m)
		} else if m*m < int64(x) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return int(r)
}
