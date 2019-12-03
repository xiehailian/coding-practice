package math

import "math"

//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数 dividend 除以除数 divisor 得到的商。
//
//示例 1:
//
//输入: dividend = 10, divisor = 3
//输出: 3
//示例 2:
//
//输入: dividend = 7, divisor = -3
//输出: -2
//说明:
//
//被除数和除数均为 32 位有符号整数。
//除数不为 0。
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/divide-two-integers

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return -1
	}

	if divisor == -1 && dividend == math.MinInt32 {
		return math.MaxInt32
	}

	negative := dividend^divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	var left = 0
	right := dividend + 1

	for left < right {
		mid := left + (right-left)/2

		if multiply(mid, divisor) > dividend {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if negative {
		return int(1 - left)
	} else {
		return int(left - 1)
	}
}

func multiply(a, b int) int {
	var res = 0
	for b != 0 {
		if b&1 != 0 {
			res += a
		}
		a <<= 1
		b >>= 1
	}
	return res
}
