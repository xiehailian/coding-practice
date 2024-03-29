package math

import (
	"sort"
)

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
//示例 2:
//
//输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum

func combinationSum(candidates []int, target int) [][]int {

	size := len(candidates)
	if size == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	var path = []int{}
	var res = [][]int{}
	dfs(candidates, 0, size, path, &res, target)
	return res
}

func dfs(candidates []int, begin int, size int, path []int, res *[][]int, target int) {

	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	for i := begin; i < size; i++ {
		residue := target - candidates[i]
		if residue < 0 {
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, i, size, path, res, residue)
		path = path[:len(path)-1]
	}
}
