package array

//给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
//
//找到所有出现两次的元素。
//
//你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
//
//示例：
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[2,3]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-all-duplicates-in-an-array

func findDuplicates(nums []int) []int {

	var res []int
	if len(nums) <= 1 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] -1 != i {
			res = append(res, nums[i])
		}
	}
	return res
}