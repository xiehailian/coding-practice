package math

import (
	"strconv"
	"strings"
)

//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//给定 n 和 k，返回第 k 个排列。
//
//说明：
//
//给定 n 的范围是 [1, 9]。
//给定 k 的范围是[1,  n!]。
//示例 1:
//
//输入: n = 3, k = 3
//输出: "213"
//示例 2:
//
//输入: n = 4, k = 9
//输出: "2314"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutation-sequence

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	var nums = make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	var used = make([]bool, n)
	for i := 0; i < n; i++ {
		used[i] = false
	}
	var path = make([]string, n)
	return recursive(nums, used, n, k, 0, &path)
}

func recursive(nums []int, used []bool, n, k, depth int, path *[]string) string {
	if depth == n {
		return strings.Join(*path, "")
	}

	counts := factorial(n - 1 - depth)

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}

		if counts < k {
			k -= counts
			continue
		}
		*path = append(*path, strconv.Itoa(nums[i]))
		used[i] = true
		return recursive(nums, used, n, k, depth+1, path)
	}
	return ""
}

func factorial(n int) int {
	var res = 1
	for n > 0 {
		res *= n
		n -= 1
	}
	return res
}
